package utils

import (
	"bytes"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "kubebuilderDemo/api/v1"
	"text/template"
)

func parseTemplate(filename string, service *v1.AppService) []byte {
	T := template.Must(template.ParseFiles("kubebuilderDemo/controllers/template/" + filename))
	wr := new(bytes.Buffer)
	if err := T.Execute(wr, service); err != nil {
		logrus.Errorf("%s  template error : %s", filename, err)
		return nil
	}

	return wr.Bytes()
}

func NewDeployment(appSvc *v1.AppService) *appsv1.Deployment {
	d := &appsv1.Deployment{}
	err := yaml.Unmarshal(parseTemplate("deployment.yml", appSvc), d)
	if err != nil {
		logrus.Errorf("New Deployment error : %s", err)

		return nil
	}
	return d
}

func NewIngress(appSvc *v1.AppService) *networkingv1.Ingress {
	i := &networkingv1.Ingress{}
	err := yaml.Unmarshal(parseTemplate("ingress.yml", appSvc), i)
	if err != nil {
		panic(err)
	}
	return i
}

func NewService(appSvc *v1.AppService) *corev1.Service {
	s := &corev1.Service{}
	err := yaml.Unmarshal(parseTemplate("service.yml", appSvc), s)
	if err != nil {
		panic(err)
	}
	return s
}
