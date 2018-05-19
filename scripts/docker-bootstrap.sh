# bootstrap a 3 node etcd cluster running on the local machine using docker
# derived from https://github.com/coreos/etcd/blob/master/Documentation/op-guide/container.md

REGISTRY=gcr.io/etcd-development/etcd
ETCD_VERSION=v3.3.5
TOKEN=token
CLUSTER_STATE=new

# due to running on the same machine, we need to change up the ports and volumes so the containers don't collide

# node 1 settings
NAME_1=etcd-node-1
PORT_1_CLIENT=7777
PORT_1_PEER=7877
VOLUME_1=etcd1-data

# node 2 settings
NAME_2=etcd-node-2
PORT_2_CLIENT=7778
PORT_2_PEER=7878
VOLUME_2=etcd2-data

# node 3 settings
NAME_3=etcd-node-3
PORT_3_CLIENT=7779
PORT_3_PEER=7879
VOLUME_3=etcd3-data

# general settings
HOST=192.168.1.73
CLUSTER=${NAME_1}=http://${HOST}:${PORT_1_PEER},${NAME_2}=http://${HOST}:${PORT_2_PEER},${NAME_3}=http://${HOST}:${PORT_3_PEER}

# create the data volumes
sudo docker volume create --name ${VOLUME_1}
sudo docker volume create --name ${VOLUME_2}
sudo docker volume create --name ${VOLUME_3}

# node 1
sudo docker run \
  -d \
  -p ${PORT_1_CLIENT}:2379 \
  -p ${PORT_1_PEER}:2380 \
  --volume=${VOLUME_1}:/etcd-data \
  --name ${NAME_1} ${REGISTRY}:${ETCD_VERSION} \
  /usr/local/bin/etcd \
  --name ${NAME_1} \
  --data-dir=/etcd-data \
  --advertise-client-urls http://${HOST}:${PORT_1_CLIENT} \
  --listen-client-urls http://0.0.0.0:2379 \
  --initial-advertise-peer-urls http://${HOST}:${PORT_1_PEER} \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-cluster ${CLUSTER} \
  --initial-cluster-state ${CLUSTER_STATE} \
  --initial-cluster-token ${TOKEN}

# node 2
sudo docker run \
  -d \
  -p ${PORT_2_CLIENT}:2379 \
  -p ${PORT_2_PEER}:2380 \
  --volume=${VOLUME_2}:/etcd-data \
  --name ${NAME_2} ${REGISTRY}:${ETCD_VERSION} \
  /usr/local/bin/etcd \
  --name ${NAME_2} \
  --data-dir=/etcd-data \
  --advertise-client-urls http://${HOST}:${PORT_2_CLIENT} \
  --listen-client-urls http://0.0.0.0:2379 \
  --initial-advertise-peer-urls http://${HOST}:${PORT_2_PEER} \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-cluster ${CLUSTER} \
  --initial-cluster-state ${CLUSTER_STATE} \
  --initial-cluster-token ${TOKEN}

# node 3
sudo docker run \
  -d \
  -p ${PORT_3_CLIENT}:2379 \
  -p ${PORT_3_PEER}:2380 \
  --volume=${VOLUME_3}:/etcd-data \
  --name ${NAME_3} ${REGISTRY}:${ETCD_VERSION} \
  /usr/local/bin/etcd \
  --name ${NAME_3} \
  --data-dir=/etcd-data \
  --advertise-client-urls http://${HOST}:${PORT_3_CLIENT} \
  --listen-client-urls http://0.0.0.0:2379 \
  --initial-advertise-peer-urls http://${HOST}:${PORT_3_PEER} \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-cluster ${CLUSTER} \
  --initial-cluster-state ${CLUSTER_STATE} \
  --initial-cluster-token ${TOKEN}