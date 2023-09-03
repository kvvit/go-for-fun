package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter the environment as argument $KUBECONFIG_DEV or $KUBECONFIG_PROD")
		return
	}
	kubeconfig := os.Args[1]
	configPath, _ := clientcmd.BuildConfigFromFlags("", kubeconfig)
	clientset, err := kubernetes.NewForConfig(configPath)
	if err != nil {
		log.Fatalf("failed to create kubernetes client: %s\n", err)
		return
	}

	ctx := context.TODO()

	nodes, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalf("failed to get list nodes: %s\n", err)
		return
	}
	for n, node := range nodes.Items {
		fmt.Printf(" Node %s has number %d\n", node.Name, n+1)
	}
	var nodeNumber int
	fmt.Println("Please enter the node number you want to check")
	fmt.Scanln(&nodeNumber)
	if nodeNumber > len(nodes.Items) {
		fmt.Println("Please enter a valid number")
		return
	}

	podHash := make(chan string)
	go func() {
		podHash <- randSeq(20)
		close(podHash)
	}()

	hashName := <-podHash
	podName := fmt.Sprintf("node-shell-%s", hashName)

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: podName, Namespace: "kube-system"},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				v1.Container{
					Name:  "shell",
					Image: "docker.io/alpine:3.17",
					SecurityContext: &v1.SecurityContext{
						Privileged: &[]bool{true}[0],
					},
					Command: []string{
						"nsenter",
					},
					Args: []string{
						"-t", "1", "-m", "-u", "-i", "-n", "sleep", "14000",
					},
				},
			},
			RestartPolicy:                 v1.RestartPolicyNever,
			HostIPC:                       true,
			HostPID:                       true,
			HostNetwork:                   true,
			NodeName:                      nodes.Items[nodeNumber-1].Name,
			EnableServiceLinks:            &[]bool{true}[0],
			TerminationGracePeriodSeconds: &[]int64{1}[0],
			Tolerations: []v1.Toleration{
				{
					Operator: v1.TolerationOpExists,
				},
			},
		},
	}

	_, err = clientset.CoreV1().Pods("kube-system").Create(ctx, pod, metav1.CreateOptions{})

	if err != nil {
		log.Fatalf("failed to create pod: %s\n", err)
	}

	fmt.Println("Awaiting the pod is up")

	// Wait for the Pod to be Running
	wait.PollImmediate(time.Second*3, time.Minute*10, func() (bool, error) {
		pod, err := clientset.CoreV1().Pods("kube-system").Get(ctx, podName, metav1.GetOptions{})
		if err != nil {
			fmt.Printf("Error getting Pod status: %v\n", err)
			return false, nil
		}
		if pod.Status.Phase == "Running" {
			return true, nil
		}
		fmt.Printf("Status still %s\n", pod.Status.Phase)
		return false, nil
	})

	fmt.Printf("Pod %s has been created\n", podName)
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
