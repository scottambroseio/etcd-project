# simple systems check script to run basic commands so we can check the output looks sane

etcdctl --endpoints=http://192.168.1.73:7777 member list
etcdctl --endpoints=http://192.168.1.73:7778 member list
etcdctl --endpoints=http://192.168.1.73:7779 member list
etcdctl --endpoints=http://192.168.1.73:7777 put foo bar
etcdctl --endpoints=http://192.168.1.73:7777 get foo
etcdctl --endpoints=http://192.168.1.73:7778 get foo
etcdctl --endpoints=http://192.168.1.73:7779 get foo