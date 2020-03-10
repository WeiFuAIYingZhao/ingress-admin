package models

//package main

import (
	//"encoding/json"
	//"errors"
	"encoding/json"
	"fmt"
	kubedata "github.com/liyinda/ingress-admin/backend/api/database"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	//"k8s.io/client-go/kubernetes"
	//"reflect"
)

//如果读取不到etcd信息返回异常?
//func (domainnames *DomainName) DnsList() (domains []DomainName, err error) {
//
//	s := `[{"id": "650000199402171139","domainname": "www.baidu.com","status": "deleted","type": "name","rdata": "10.10.10.10","ttl": "4316"},{"id": "650000199402171131","domainname": "sdfsdf","status": "deleted","type": "name","rdata": "2017-03-11 05:35:45","ttl": "4317"}]`
//	//s, _ := orm.GetEtcdPrefix()
//
//	err = json.Unmarshal([]byte(s), &domains)
//	if err != nil {
//		return
//	}
//	return
//
//}

type IngressMeta struct {
	IngressName string
	NameSpace   string
	Host        string
	ServiceName string
	ServicePort intstr.IntOrString
}

type IngressJson struct {
	IngressMeta string
}

//var clientset *kubernetes.Clientset

//在namespace中获取该命名空间中的ingress信息列表
//func (ingmeta *IngressMeta) ListIngress(clientset *kubernetes.Clientset, namespace string) v1beta1.Ingress {
func ListIngress(namespace string) (value string, err error) {
	//if clientset != nil {
	//	clientset, _ := kubeconfig.NewClientset()
	//}

	clientset, _ := kubedata.NewClientset()
	//defer clientset.Close()

	//创建map和slince
	var m map[string]interface{}
	var s []map[string]interface{}

	//列出namespace下的所有ingress
	ingressList, err := clientset.ExtensionsV1beta1().Ingresses(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	ingressCtrls := ingressList.Items

	if len(ingressCtrls) > 0 {
		for _, ingress := range ingressCtrls {
			//fmt.Printf("Rules %s exists in namespace %s\n", ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServiceName, ingress.Namespace)
			//将获取的ingress中的信息JSON序列化
			m = make(map[string]interface{})
			m["id"] = "650000199402171139"
			m["ingressname"] = ingress.Name
			m["namespace"] = namespace
			m["host"] = ingress.Spec.Rules[0].Host
			m["servicename"] = ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServiceName
			m["serviceport"] = ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServicePort
			s = append(s, m)
			data, err := json.Marshal(s)
			if err != nil {
				fmt.Printf("json error %v\n", err)
			}
			value = string(data)
			//fmt.Println(string(data))
		}
	} else {
		fmt.Println("no ingress found")
	}
	return value, err
}

//func main() {
//	//func (ingress *Ingress) GetList() v1beta1.Ingress {
//	clientset, _ := kubeconfig.NewClientset()
//	list, _ := ListIngress(clientset, "kube-system")
//	fmt.Println("type:", reflect.TypeOf(list))
//	fmt.Println(list)
//	//return list
//}
