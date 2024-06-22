package bb

import "fmt"

func request(reqType, reqCode int) int {
	return (reqType << 8) | reqCode
}

func generateVariable(name string, value int, comment string) string {
	if comment == "" {
		return fmt.Sprintf("const %s = 0x%08x", name, value)
	}
	return fmt.Sprintf("const %s = 0x%08x /** %s */", name, value, comment)
}

func GenerateCfg() []string {
	ss := make([]string, 0)
	ss = append(ss, generateVariable("BB_CFG_AP_BASIC", request(BB_REQ_CFG, 0), "AP角色的基本配置命令字"))
	ss = append(ss, generateVariable("BB_CFG_DEV_BASIC", request(BB_REQ_CFG, 1), "DEV角色的基本配置命令字"))
	ss = append(ss, generateVariable("BB_CFG_CHANNEL", request(BB_REQ_CFG, 2), "信道配置命令字"))
	ss = append(ss, generateVariable("BB_CFG_CANDIDATES", request(BB_REQ_CFG, 3), "AP候选人配置命令字"))
	ss = append(ss, generateVariable("BB_CFG_USER_PARA", request(BB_REQ_CFG, 4), "基带用户参数配置命令字"))
	ss = append(ss, generateVariable("BB_CFG_SLOT_RX_MCS", request(BB_REQ_CFG, 5), "基于SLOT的MCS策略配置命令字"))
	ss = append(ss, generateVariable("BB_CFG_ANY_CHANNEL", request(BB_REQ_CFG, 6), "任意工作信道配置命令字，依赖于功能宏BB_CONFIG_ENABLE_ANY_CHANNEL"))
	ss = append(ss, generateVariable("BB_CFG_DISTC", request(BB_REQ_CFG, 7), "基带测距功能命令字"))
	ss = append(ss, generateVariable("BB_CFG_AP_SYNC_MODE", request(BB_REQ_CFG, 9), "配置AP同步模式，默认为非同步模式"))
	ss = append(ss, generateVariable("BB_CFG_BR_HOP_POLICY", request(BB_REQ_CFG, 10), "配置BR的跳频策略"))
	ss = append(ss, generateVariable("BB_CFG_PWR_BASIC", request(BB_REQ_CFG, 11), "配置功率参数"))
	ss = append(ss, generateVariable("BB_CFG_RC_HOP_POLICY", request(BB_REQ_CFG, 12), "配置选择性跳频策略"))
	ss = append(ss, generateVariable("BB_CFG_POWER_SAVE", request(BB_REQ_CFG, 13), "配置节能模式"))
	return ss
}

func GenerateGet() []string {
	ss := make([]string, 0)
	ss = append(ss, generateVariable("BB_GET_STATUS", request(BB_REQ_GET, 0), "读取基带工作状态命令字"))
	ss = append(ss, generateVariable("BB_GET_PAIR_RESULT", request(BB_REQ_GET, 1), "读取配对命令结果命令字"))
	ss = append(ss, generateVariable("BB_GET_AP_MAC", request(BB_REQ_GET, 2), "DEV读取目标AP的MAC命令字"))
	ss = append(ss, generateVariable("BB_GET_CANDIDATES", request(BB_REQ_GET, 3), "AP读取指定SLOT的候选人命令字"))
	ss = append(ss, generateVariable("BB_GET_USER_QUALITY", request(BB_REQ_GET, 4), "读取物理用户信号质量"))
	ss = append(ss, generateVariable("BB_GET_DISTC_RESULT", request(BB_REQ_GET, 5), "读取测距结果"))
	ss = append(ss, generateVariable("BB_GET_MCS", request(BB_REQ_GET, 6), "读取指定SLOT的MCS及理论吞吐率"))
	ss = append(ss, generateVariable("BB_GET_POWER_MODE", request(BB_REQ_GET, 7), "读取功率开、闭环模式"))
	ss = append(ss, generateVariable("BB_GET_CUR_POWER", request(BB_REQ_GET, 8), "读取当前发射功率值"))
	ss = append(ss, generateVariable("BB_GET_POWER_AUTO", request(BB_REQ_GET, 9), "当前功率自适应是否开启"))
	ss = append(ss, generateVariable("BB_GET_CHAN_INFO", request(BB_REQ_GET, 10), "获取信道相关信息，如信道列表、周期扫频结果等"))
	ss = append(ss, generateVariable("BB_GET_PEER_QUALITY", request(BB_REQ_GET, 11), "获取对端信号质量（数据通道）"))
	ss = append(ss, generateVariable("BB_GET_AP_TIME", request(BB_REQ_GET, 12), "获取AP时间戳"))
	ss = append(ss, generateVariable("BB_GET_BAND_INFO", request(BB_REQ_GET, 13), "获取频段信息"))
	ss = append(ss, generateVariable("BB_GET_REMOTE", request(BB_REQ_GET, 14), "获取通讯对端的配置"))
	ss = append(ss, generateVariable("BB_GET_REG", request(BB_REQ_GET, 100), "基带寄存器读取命令字，本类型用于调试诊断"))
	ss = append(ss, generateVariable("BB_GET_CFG", request(BB_REQ_GET, 101), "基带配置文件读取命令字"))
	ss = append(ss, generateVariable("BB_GET_DBG_MODE", request(BB_REQ_GET, 102), "获取SDK debug模式状态"))
	ss = append(ss, generateVariable("BB_GET_HARDWARE_VERSION", request(BB_REQ_GET, 103), "软硬件版本"))
	ss = append(ss, generateVariable("BB_GET_FIRMWARE_VERSION", request(BB_REQ_GET, 104), "固件版本"))
	ss = append(ss, generateVariable("BB_GET_SYS_INFO", request(BB_REQ_GET, 105), "获取系统信息，如运行时间，软硬件版本等"))
	ss = append(ss, generateVariable("BB_GET_USER_INFO", request(BB_REQ_GET, 106), "获取基带用户基带信息"))
	ss = append(ss, generateVariable("BB_GET_1V1_INFO", request(BB_REQ_GET, 107), "获取基带用户基带信息"))
	ss = append(ss, generateVariable("BB_GET_PRJ_DISPATCH", request(BB_REQ_GET, 200), "二级GET命令分发"))
	return ss
}

func GenerateSet() []string {
	ss := make([]string, 0)
	ss = append(ss, generateVariable("BB_SET_EVENT_SUBSCRIBE", request(BB_REQ_SET, 0), "事件订阅类型命令字"))
	ss = append(ss, generateVariable("BB_SET_EVENT_UNSUBSCRIBE", request(BB_REQ_SET, 1), "事件反订阅类型命令字"))
	ss = append(ss, generateVariable("BB_SET_PAIR_MODE", request(BB_REQ_SET, 2), "设置指定SLOT进入配对模式命令字"))
	ss = append(ss, generateVariable("BB_SET_AP_MAC", request(BB_REQ_SET, 3), "DEV设置AP的MAC命令字"))
	ss = append(ss, generateVariable("BB_SET_CANDIDATES", request(BB_REQ_SET, 4), "AP设置候选人命令字"))
	ss = append(ss, generateVariable("BB_SET_CHAN_MODE", request(BB_REQ_SET, 5), "设置信道工作模式"))
	ss = append(ss, generateVariable("BB_SET_CHAN", request(BB_REQ_SET, 6), "设置信道"))
	ss = append(ss, generateVariable("BB_SET_POWER_MODE", request(BB_REQ_SET, 7), "设置功率开、闭环模式"))
	ss = append(ss, generateVariable("BB_SET_POWER", request(BB_REQ_SET, 8), "设置发射功率值"))
	ss = append(ss, generateVariable("BB_SET_POWER_AUTO", request(BB_REQ_SET, 9), "使能功率自适应"))
	ss = append(ss, generateVariable("BB_SET_HOT_UPGRADE_WRITE", request(BB_REQ_SET, 10), "基带热升级命令"))
	ss = append(ss, generateVariable("BB_SET_HOT_UPGRADE_CRC32", request(BB_REQ_SET, 11), "基带热升级校验命令"))
	ss = append(ss, generateVariable("BB_SET_MCS_MODE", request(BB_REQ_SET, 12), "设置MCS控制模式手动、自动"))
	ss = append(ss, generateVariable("BB_SET_MCS", request(BB_REQ_SET, 13), "设置MCS挡位，仅在手动模式下支持"))
	ss = append(ss, generateVariable("BB_SET_SYS_REBOOT", request(BB_REQ_SET, 14), "系统重启"))
	ss = append(ss, generateVariable("BB_SET_MASTER_DEV", request(BB_REQ_SET, 15), "设置导演模式下的主DEV设备，仅导演模式AP侧支持"))
	ss = append(ss, generateVariable("BB_SET_FRAME_CHANGE", request(BB_REQ_SET, 16), "运行中帧结构改变(仅1V1模式)"))
	ss = append(ss, generateVariable("BB_SET_COMPLIANCE_MODE", request(BB_REQ_SET, 17), "设置频点合规模式"))
	ss = append(ss, generateVariable("BB_SET_BAND_MODE", request(BB_REQ_SET, 18), "设置频段切换模式"))
	ss = append(ss, generateVariable("BB_SET_BAND", request(BB_REQ_SET, 19), "设置工作频段"))
	ss = append(ss, generateVariable("BB_FORCE_CLS_SOCKET_ALL", request(BB_REQ_SET, 20), "强制关闭所有socket 可能会造成socket信息不同步"))
	ss = append(ss, generateVariable("BB_SET_REMOTE", request(BB_REQ_SET, 21), "设置通讯对端的配置"))
	ss = append(ss, generateVariable("BB_SET_BANDWIDTH", request(BB_REQ_SET, 22), "1V1模式下手动改变bandwidth"))
	ss = append(ss, generateVariable("BB_SET_DFS", request(BB_REQ_SET, 23), "设置DFS检测"))
	ss = append(ss, generateVariable("BB_SET_RF", request(BB_REQ_SET, 24), "RF收发开关动态控制"))
	ss = append(ss, generateVariable("BB_SET_POWER_SAVE_MODE", request(BB_REQ_SET, 25), "1V1模式低功耗策略模式"))
	ss = append(ss, generateVariable("BB_SET_POWER_SAVE", request(BB_REQ_SET, 26), "1V1模式低功耗手动周期"))
	ss = append(ss, generateVariable("BB_SET_REG", request(BB_REQ_SET, 100), "基带寄存器写入命令字，本类型用于调试诊断"))
	ss = append(ss, generateVariable("BB_SET_CFG", request(BB_REQ_SET, 101), "基带配置文件写入命令字"))
	ss = append(ss, generateVariable("BB_RESET_CFG", request(BB_REQ_SET, 102), "reset基带配置文件命令字"))
	ss = append(ss, generateVariable("BB_SET_PLOT", request(BB_REQ_SET, 103), "设置基带plot debug参数"))
	ss = append(ss, generateVariable("BB_SET_DBG_MODE", request(BB_REQ_SET, 104), "设置基带SDK进入debug mode（软件不工作）"))
	ss = append(ss, generateVariable("BB_SET_FREQ", request(BB_REQ_SET, 105), "设置物理用户的发送或接收频率"))
	ss = append(ss, generateVariable("BB_SET_TX_MCS", request(BB_REQ_SET, 106), "设置物理用户的发送MCS"))
	ss = append(ss, generateVariable("BB_SET_RESET", request(BB_REQ_SET, 107), "基带RESET命令"))
	ss = append(ss, generateVariable("BB_SET_TX_PATH", request(BB_REQ_SET, 108), "设置无线发射通道，用于功率测试"))
	ss = append(ss, generateVariable("BB_SET_RX_PATH", request(BB_REQ_SET, 109), "设置无线接收通道，用于灵敏度测试"))
	ss = append(ss, generateVariable("BB_SET_POWER_OFFSET", request(BB_REQ_SET, 110), "设置功率补偿，用于功率测试"))
	ss = append(ss, generateVariable("BB_SET_POWER_TEST_MODE", request(BB_REQ_SET, 111), "进入产测功率测试模式"))
	ss = append(ss, generateVariable("BB_SET_SENSE_TEST_MODE", request(BB_REQ_SET, 112), "进入产测灵敏度测试模式"))
	ss = append(ss, generateVariable("BB_SET_ORIG_CFG", request(BB_REQ_SET, 113), "加载原始镜像配置，仅在基带IDLE状态支持"))
	ss = append(ss, generateVariable("BB_SET_SINGLE_TONE", request(BB_REQ_SET, 114), "单音信号，需要先进入debug模式"))
	ss = append(ss, generateVariable("BB_SET_PURE_SLOT", request(BB_REQ_SET, 115), "纯图传模式，注意暂不支持退出"))
	ss = append(ss, generateVariable("BB_SET_FACTORY_POWER_OFFSET_SAVE", request(BB_REQ_SET, 116), "产测功率校准保存指令"))
	ss = append(ss, generateVariable("BB_SET_PRJ_DISPATCH", request(BB_REQ_SET, 200), "二级SET命令分发"))
	return ss
}

func GenerateSpecialControl() []string {
	ss := make([]string, 0)
	ss = append(ss, generateVariable("BB_START_REQ", request(BB_REQ_RPC_IOCTL, 0), ""))
	ss = append(ss, generateVariable("BB_STOP_REQ", request(BB_REQ_RPC_IOCTL, 1), ""))
	ss = append(ss, generateVariable("BB_INIT_REQ", request(BB_REQ_RPC_IOCTL, 2), ""))
	ss = append(ss, generateVariable("BB_DEINIT_REQ", request(BB_REQ_RPC_IOCTL, 3), ""))
	ss = append(ss, generateVariable("BB_RPC_GET_LIST", request(BB_REQ_RPC, 0), "获取8030可用列表"))
	ss = append(ss, generateVariable("BB_RPC_SEL_ID", request(BB_REQ_RPC, 1), "选择指定的8030进行通讯"))
	ss = append(ss, generateVariable("BB_RPC_GET_MAC", request(BB_REQ_RPC, 2), "获取指定8030的mac"))
	ss = append(ss, generateVariable("BB_RPC_GET_HOTPLUG_EVENT", request(BB_REQ_RPC, 3), "获取设备上下线通知"))
	ss = append(ss, generateVariable("BB_RPC_SOCK_BUF_STA", request(BB_REQ_RPC, 4), "查询socket buff状态"))
	ss = append(ss, generateVariable("BB_RPC_TEST", request(BB_REQ_RPC, 5), "测试服务器连通性"))
	ss = append(ss, generateVariable("BB_RPC_SERIAL_LIST", request(BB_REQ_PLAT_CTL, 0), "获取串口列表"))
	ss = append(ss, generateVariable("BB_RPC_SERIAL_SETUP", request(BB_REQ_PLAT_CTL, 1), "设置串口"))
	return ss
}

// PrintConstants prints the constants
func PrintConstants() {
	printWithName := func(name string, content []string) {
		s := fmt.Sprintf("// %s", name)
		fmt.Println(s)
		for _, s := range content {
			fmt.Println(s)
		}
		fmt.Print("\n")
	}
	printWithName("cfg", GenerateCfg())
	printWithName("get", GenerateGet())
	printWithName("set", GenerateSet())
	printWithName("special control", GenerateSpecialControl())
}
