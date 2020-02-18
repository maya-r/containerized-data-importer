/*
 * This file is part of the CDI project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2020 Red Hat, Inc.
 *
 */

package webhooks

import (
	//"encoding/json"
	//"fmt"
	//"net/url"
	//"reflect"

	"k8s.io/api/admission/v1beta1"
	//v1 "k8s.io/api/core/v1"
	//k8serrors "k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//k8sfield "k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"

	//cdicorev1alpha1 "kubevirt.io/containerized-data-importer/pkg/apis/core/v1alpha1"
	//"kubevirt.io/containerized-data-importer/pkg/controller"
)

type cdiValidatingWebhook struct {
	client kubernetes.Interface
}

func (wh *cdiValidatingWebhook) Admit(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	//raw := ar.Request.Object.Raw
	//dv := cdicorev1alpha1.DataVolume{}

	//err := json.Unmarshal(raw, &dv)
	//if err != nil {
	//	return toAdmissionResponseError(err)
	//}

	klog.Infof("Got a request something something XXXXXXXXXXXXX")
	if ar.Request.Operation == v1beta1.Delete {
		klog.Infof("XXXXXXXXXXXXXX admission webhook delete")
		//causes = append(causes, metav1.StatusCause{
		//	Type:    metav1.CauseTypeFieldValueDuplicate,
		//	Message: fmt.Sprintf("Destination PVC already exists"),
		//	Field:   k8sfield.NewPath("DataVolume").Child("Name").String(),
		//})
		//return toRejectedAdmissionResponse(causes)
	}

	reviewResponse := v1beta1.AdmissionResponse{}
	reviewResponse.Allowed = true
	return &reviewResponse
}
