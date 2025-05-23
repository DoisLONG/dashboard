package service

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

func GetServiceList(client k8sClient.Interface, namespace string) (*corev1.ServiceList, error) {
	klog.V(4).Infof("Getting service list")

	list, err := client.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		klog.Errorf("Failed to get service list: %v", err)
		return nil, err
	}

	return list, nil
}

func GetService(client k8sClient.Interface, namespace, name string) (*corev1.Service, error) {
	klog.V(4).Infof("Getting service %s in namespace %s", name, namespace)

	service, err := client.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		klog.Errorf("Failed to get service %s in namespace %s: %v", name, namespace, err)
		return nil, err
	}

	return service, nil
}

func CreateService(client k8sClient.Interface, namespace string, service *corev1.Service) (*corev1.Service, error) {
	klog.V(4).Infof("Creating service %s in namespace %s", service.Name, namespace)

	createdService, err := client.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		klog.Errorf("Failed to create service %s in namespace %s: %v", service.Name, namespace, err)
		return nil, err
	}

	return createdService, nil
}

func UpdateService(client k8sClient.Interface, namespace string, service *corev1.Service) (*corev1.Service, error) {
	klog.V(4).Infof("Updating service %s in namespace %s", service.Name, namespace)

	updatedService, err := client.CoreV1().Services(namespace).Update(context.TODO(), service, metav1.UpdateOptions{})
	if err != nil {
		klog.Errorf("Failed to update service %s in namespace %s: %v", service.Name, namespace, err)
		return nil, err
	}

	return updatedService, nil
}

func DeleteService(client k8sClient.Interface, namespace, name string) error {
	klog.V(4).Infof("Deleting service %s in namespace %s", name, namespace)

	err := client.CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		klog.Errorf("Failed to delete service %s in namespace %s: %v", name, namespace, err)
		return err
	}

	return nil
}
