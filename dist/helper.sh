#!/bin/bash

set -x

nam=$1;
con=$2;
pod=$3;
shift;
shift;
shift;
shift;

#kubectl exec -i -n $nam -c $con $pod -- "rsync --server -vlogDtpr . /tmp"
kubectl exec -i -n $nam -c $con $pod -- "rsync" "--server" "-vlogDtpre.iLsfxCIvu" "--stats" "." "/tmp"
