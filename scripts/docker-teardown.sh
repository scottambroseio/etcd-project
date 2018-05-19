# tear down script to nuke the previously bootstrapped cluster - useful for a clean start (it also nukes the data volumes)

sudo docker kill etcd-node-1
sudo docker rm etcd-node-1
sudo docker volume rm etcd1-data
sudo docker kill etcd-node-2
sudo docker rm etcd-node-2
sudo docker volume rm etcd2-data
sudo docker kill etcd-node-3
sudo docker rm etcd-node-3
sudo docker volume rm etcd3-data