// Code generated by CodeGenerator. Not generate if exist
//
// Source: account.proto
// Time: 2022-11-30 15:50:38

package config

type RpcServerBaseConfig struct {
	ListenOn string
	RpcMode  string `json:",default=grpc,options=grpc|rabbit|inter"`
}

type Config struct {
	RpcServerBaseConfig
	Spec accountSpecialConfig
}

type accountSpecialConfig struct {
}