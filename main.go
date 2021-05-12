package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	ilsaclient "github.com/jungmin17/ilsa-clientset-example/pkg/client/clientset/versioned"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()


	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	//k8s client setting
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ilsaClientset, err := ilsaclient.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//something....
	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, item := range namespaceList.Items {
		ilsas, err := ilsaClientset.IlsaV1().MyPods(item.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		for _, ilsa := range ilsas.Items {
			fmt.Println("namespace :", item.Name)
			fmt.Println("name :", ilsa.Name)
			fmt.Println("replica :",ilsa.Spec.Replicas)
		}

	}
	fmt.Println("------")
}
