

kubectl run mysql-client --image=mysql:8 -it --rm --restart=Never -- bash

kubectl exec -it mysql-client -- bash
