package bb

// Generated content
// modify with care

// cfg
const BB_CFG_AP_BASIC RequestId = 0x00000000      /** AP角色的基本配置命令字 */
const BB_CFG_DEV_BASIC RequestId = 0x00000001     /** DEV角色的基本配置命令字 */
const BB_CFG_CHANNEL RequestId = 0x00000002       /** 信道配置命令字 */
const BB_CFG_CANDIDATES RequestId = 0x00000003    /** AP候选人配置命令字 */
const BB_CFG_USER_PARA RequestId = 0x00000004     /** 基带用户参数配置命令字 */
const BB_CFG_SLOT_RX_MCS RequestId = 0x00000005   /** 基于SLOT的MCS策略配置命令字 */
const BB_CFG_ANY_CHANNEL RequestId = 0x00000006   /** 任意工作信道配置命令字，依赖于功能宏BB_CONFIG_ENABLE_ANY_CHANNEL */
const BB_CFG_DISTC RequestId = 0x00000007         /** 基带测距功能命令字 */
const BB_CFG_AP_SYNC_MODE RequestId = 0x00000009  /** 配置AP同步模式，默认为非同步模式 */
const BB_CFG_BR_HOP_POLICY RequestId = 0x0000000a /** 配置BR的跳频策略 */
const BB_CFG_PWR_BASIC RequestId = 0x0000000b     /** 配置功率参数 */
const BB_CFG_RC_HOP_POLICY RequestId = 0x0000000c /** 配置选择性跳频策略 */
const BB_CFG_POWER_SAVE RequestId = 0x0000000d    /** 配置节能模式 */

// get
const BB_GET_STATUS RequestId = 0x01000000           /** 读取基带工作状态命令字 */
const BB_GET_PAIR_RESULT RequestId = 0x01000001      /** 读取配对命令结果命令字 */
const BB_GET_AP_MAC RequestId = 0x01000002           /** DEV读取目标AP的MAC命令字 */
const BB_GET_CANDIDATES RequestId = 0x01000003       /** AP读取指定SLOT的候选人命令字 */
const BB_GET_USER_QUALITY RequestId = 0x01000004     /** 读取物理用户信号质量 */
const BB_GET_DISTC_RESULT RequestId = 0x01000005     /** 读取测距结果 */
const BB_GET_MCS RequestId = 0x01000006              /** 读取指定SLOT的MCS及理论吞吐率 */
const BB_GET_POWER_MODE RequestId = 0x01000007       /** 读取功率开、闭环模式 */
const BB_GET_CUR_POWER RequestId = 0x01000008        /** 读取当前发射功率值 */
const BB_GET_POWER_AUTO RequestId = 0x01000009       /** 当前功率自适应是否开启 */
const BB_GET_CHAN_INFO RequestId = 0x0100000a        /** 获取信道相关信息，如信道列表、周期扫频结果等 */
const BB_GET_PEER_QUALITY RequestId = 0x0100000b     /** 获取对端信号质量（数据通道） */
const BB_GET_AP_TIME RequestId = 0x0100000c          /** 获取AP时间戳 */
const BB_GET_BAND_INFO RequestId = 0x0100000d        /** 获取频段信息 */
const BB_GET_REMOTE RequestId = 0x0100000e           /** 获取通讯对端的配置 */
const BB_GET_REG RequestId = 0x01000064              /** 基带寄存器读取命令字，本类型用于调试诊断 */
const BB_GET_CFG RequestId = 0x01000065              /** 基带配置文件读取命令字 */
const BB_GET_DBG_MODE RequestId = 0x01000066         /** 获取SDK debug模式状态 */
const BB_GET_HARDWARE_VERSION RequestId = 0x01000067 /** 软硬件版本 */
const BB_GET_FIRMWARE_VERSION RequestId = 0x01000068 /** 固件版本 */
const BB_GET_SYS_INFO RequestId = 0x01000069         /** 获取系统信息，如运行时间，软硬件版本等 */
const BB_GET_USER_INFO RequestId = 0x0100006a        /** 获取基带用户基带信息 */
const BB_GET_1V1_INFO RequestId = 0x0100006b         /** 获取基带用户基带信息 */
const BB_GET_PRJ_DISPATCH RequestId = 0x010000c8     /** 二级GET命令分发 */

// set
const BB_SET_EVENT_SUBSCRIBE RequestId = 0x02000000           /** 事件订阅类型命令字 */
const BB_SET_EVENT_UNSUBSCRIBE RequestId = 0x02000001         /** 事件反订阅类型命令字 */
const BB_SET_PAIR_MODE RequestId = 0x02000002                 /** 设置指定SLOT进入配对模式命令字 */
const BB_SET_AP_MAC RequestId = 0x02000003                    /** DEV设置AP的MAC命令字 */
const BB_SET_CANDIDATES RequestId = 0x02000004                /** AP设置候选人命令字 */
const BB_SET_CHAN_MODE RequestId = 0x02000005                 /** 设置信道工作模式 */
const BB_SET_CHAN RequestId = 0x02000006                      /** 设置信道 */
const BB_SET_POWER_MODE RequestId = 0x02000007                /** 设置功率开、闭环模式 */
const BB_SET_POWER RequestId = 0x02000008                     /** 设置发射功率值 */
const BB_SET_POWER_AUTO RequestId = 0x02000009                /** 使能功率自适应 */
const BB_SET_HOT_UPGRADE_WRITE RequestId = 0x0200000a         /** 基带热升级命令 */
const BB_SET_HOT_UPGRADE_CRC32 RequestId = 0x0200000b         /** 基带热升级校验命令 */
const BB_SET_MCS_MODE RequestId = 0x0200000c                  /** 设置MCS控制模式手动、自动 */
const BB_SET_MCS RequestId = 0x0200000d                       /** 设置MCS挡位，仅在手动模式下支持 */
const BB_SET_SYS_REBOOT RequestId = 0x0200000e                /** 系统重启 */
const BB_SET_MASTER_DEV RequestId = 0x0200000f                /** 设置导演模式下的主DEV设备，仅导演模式AP侧支持 */
const BB_SET_FRAME_CHANGE RequestId = 0x02000010              /** 运行中帧结构改变(仅1V1模式) */
const BB_SET_COMPLIANCE_MODE RequestId = 0x02000011           /** 设置频点合规模式 */
const BB_SET_BAND_MODE RequestId = 0x02000012                 /** 设置频段切换模式 */
const BB_SET_BAND RequestId = 0x02000013                      /** 设置工作频段 */
const BB_FORCE_CLS_SOCKET_ALL RequestId = 0x02000014          /** 强制关闭所有socket 可能会造成socket信息不同步 */
const BB_SET_REMOTE RequestId = 0x02000015                    /** 设置通讯对端的配置 */
const BB_SET_BANDWIDTH RequestId = 0x02000016                 /** 1V1模式下手动改变bandwidth */
const BB_SET_DFS RequestId = 0x02000017                       /** 设置DFS检测 */
const BB_SET_RF RequestId = 0x02000018                        /** RF收发开关动态控制 */
const BB_SET_POWER_SAVE_MODE RequestId = 0x02000019           /** 1V1模式低功耗策略模式 */
const BB_SET_POWER_SAVE RequestId = 0x0200001a                /** 1V1模式低功耗手动周期 */
const BB_SET_REG RequestId = 0x02000064                       /** 基带寄存器写入命令字，本类型用于调试诊断 */
const BB_SET_CFG RequestId = 0x02000065                       /** 基带配置文件写入命令字 */
const BB_RESET_CFG RequestId = 0x02000066                     /** reset基带配置文件命令字 */
const BB_SET_PLOT RequestId = 0x02000067                      /** 设置基带plot debug参数 */
const BB_SET_DBG_MODE RequestId = 0x02000068                  /** 设置基带SDK进入debug mode（软件不工作） */
const BB_SET_FREQ RequestId = 0x02000069                      /** 设置物理用户的发送或接收频率 */
const BB_SET_TX_MCS RequestId = 0x0200006a                    /** 设置物理用户的发送MCS */
const BB_SET_RESET RequestId = 0x0200006b                     /** 基带RESET命令 */
const BB_SET_TX_PATH RequestId = 0x0200006c                   /** 设置无线发射通道，用于功率测试 */
const BB_SET_RX_PATH RequestId = 0x0200006d                   /** 设置无线接收通道，用于灵敏度测试 */
const BB_SET_POWER_OFFSET RequestId = 0x0200006e              /** 设置功率补偿，用于功率测试 */
const BB_SET_POWER_TEST_MODE RequestId = 0x0200006f           /** 进入产测功率测试模式 */
const BB_SET_SENSE_TEST_MODE RequestId = 0x02000070           /** 进入产测灵敏度测试模式 */
const BB_SET_ORIG_CFG RequestId = 0x02000071                  /** 加载原始镜像配置，仅在基带IDLE状态支持 */
const BB_SET_SINGLE_TONE RequestId = 0x02000072               /** 单音信号，需要先进入debug模式 */
const BB_SET_PURE_SLOT RequestId = 0x02000073                 /** 纯图传模式，注意暂不支持退出 */
const BB_SET_FACTORY_POWER_OFFSET_SAVE RequestId = 0x02000074 /** 产测功率校准保存指令 */
const BB_SET_PRJ_DISPATCH RequestId = 0x020000c8              /** 二级SET命令分发 */

// special control
const BB_START_REQ RequestId = 0x0b000000
const BB_STOP_REQ RequestId = 0x0b000001
const BB_INIT_REQ RequestId = 0x0b000002
const BB_DEINIT_REQ RequestId = 0x0b000003
const BB_RPC_GET_LIST RequestId = 0x0a000000          /** 获取8030可用列表 */
const BB_RPC_SEL_ID RequestId = 0x0a000001            /** 选择指定的8030进行通讯 */
const BB_RPC_GET_MAC RequestId = 0x0a000002           /** 获取指定8030的mac */
const BB_RPC_GET_HOTPLUG_EVENT RequestId = 0x0a000003 /** 获取设备上下线通知 */
const BB_RPC_SOCK_BUF_STA RequestId = 0x0a000004      /** 查询socket buff状态 */
const BB_RPC_TEST RequestId = 0x0a000005              /** 测试服务器连通性 */
const BB_RPC_SERIAL_LIST RequestId = 0x0c000000       /** 获取串口列表 */
const BB_RPC_SERIAL_SETUP RequestId = 0x0c000001      /** 设置串口 */
