# promadapt
prometheus adapeter test setup

intended for prometheus installed with https://github.com/coreos/kube-prometheus

get --raw /apis/custom.metrics.k8s.io/v1beta1/namespaces/default/pods/*/cntx | jq
