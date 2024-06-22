package bb

// Generated content
// modify with care

// cfg
const BB_CFG_AP_BASIC = 0x00000000      /** AP角色的基本配置命令字 */
const BB_CFG_DEV_BASIC = 0x00000001     /** DEV角色的基本配置命令字 */
const BB_CFG_CHANNEL = 0x00000002       /** 信道配置命令字 */
const BB_CFG_CANDIDATES = 0x00000003    /** AP候选人配置命令字 */
const BB_CFG_USER_PARA = 0x00000004     /** 基带用户参数配置命令字 */
const BB_CFG_SLOT_RX_MCS = 0x00000005   /** 基于SLOT的MCS策略配置命令字 */
const BB_CFG_ANY_CHANNEL = 0x00000006   /** 任意工作信道配置命令字，依赖于功能宏BB_CONFIG_ENABLE_ANY_CHANNEL */
const BB_CFG_DISTC = 0x00000007         /** 基带测距功能命令字 */
const BB_CFG_AP_SYNC_MODE = 0x00000009  /** 配置AP同步模式，默认为非同步模式 */
const BB_CFG_BR_HOP_POLICY = 0x0000000a /** 配置BR的跳频策略 */
const BB_CFG_PWR_BASIC = 0x0000000b     /** 配置功率参数 */
const BB_CFG_RC_HOP_POLICY = 0x0000000c /** 配置选择性跳频策略 */
const BB_CFG_POWER_SAVE = 0x0000000d    /** 配置节能模式 */

// get
const BB_GET_STATUS = 0x00000100           /** 读取基带工作状态命令字 */
const BB_GET_PAIR_RESULT = 0x00000101      /** 读取配对命令结果命令字 */
const BB_GET_AP_MAC = 0x00000102           /** DEV读取目标AP的MAC命令字 */
const BB_GET_CANDIDATES = 0x00000103       /** AP读取指定SLOT的候选人命令字 */
const BB_GET_USER_QUALITY = 0x00000104     /** 读取物理用户信号质量 */
const BB_GET_DISTC_RESULT = 0x00000105     /** 读取测距结果 */
const BB_GET_MCS = 0x00000106              /** 读取指定SLOT的MCS及理论吞吐率 */
const BB_GET_POWER_MODE = 0x00000107       /** 读取功率开、闭环模式 */
const BB_GET_CUR_POWER = 0x00000108        /** 读取当前发射功率值 */
const BB_GET_POWER_AUTO = 0x00000109       /** 当前功率自适应是否开启 */
const BB_GET_CHAN_INFO = 0x0000010a        /** 获取信道相关信息，如信道列表、周期扫频结果等 */
const BB_GET_PEER_QUALITY = 0x0000010b     /** 获取对端信号质量（数据通道） */
const BB_GET_AP_TIME = 0x0000010c          /** 获取AP时间戳 */
const BB_GET_BAND_INFO = 0x0000010d        /** 获取频段信息 */
const BB_GET_REMOTE = 0x0000010e           /** 获取通讯对端的配置 */
const BB_GET_REG = 0x00000164              /** 基带寄存器读取命令字，本类型用于调试诊断 */
const BB_GET_CFG = 0x00000165              /** 基带配置文件读取命令字 */
const BB_GET_DBG_MODE = 0x00000166         /** 获取SDK debug模式状态 */
const BB_GET_HARDWARE_VERSION = 0x00000167 /** 软硬件版本 */
const BB_GET_FIRMWARE_VERSION = 0x00000168 /** 固件版本 */
const BB_GET_SYS_INFO = 0x00000169         /** 获取系统信息，如运行时间，软硬件版本等 */
const BB_GET_USER_INFO = 0x0000016a        /** 获取基带用户基带信息 */
const BB_GET_1V1_INFO = 0x0000016b         /** 获取基带用户基带信息 */
const BB_GET_PRJ_DISPATCH = 0x000001c8     /** 二级GET命令分发 */

// set
const BB_SET_EVENT_SUBSCRIBE = 0x00000200           /** 事件订阅类型命令字 */
const BB_SET_EVENT_UNSUBSCRIBE = 0x00000201         /** 事件反订阅类型命令字 */
const BB_SET_PAIR_MODE = 0x00000202                 /** 设置指定SLOT进入配对模式命令字 */
const BB_SET_AP_MAC = 0x00000203                    /** DEV设置AP的MAC命令字 */
const BB_SET_CANDIDATES = 0x00000204                /** AP设置候选人命令字 */
const BB_SET_CHAN_MODE = 0x00000205                 /** 设置信道工作模式 */
const BB_SET_CHAN = 0x00000206                      /** 设置信道 */
const BB_SET_POWER_MODE = 0x00000207                /** 设置功率开、闭环模式 */
const BB_SET_POWER = 0x00000208                     /** 设置发射功率值 */
const BB_SET_POWER_AUTO = 0x00000209                /** 使能功率自适应 */
const BB_SET_HOT_UPGRADE_WRITE = 0x0000020a         /** 基带热升级命令 */
const BB_SET_HOT_UPGRADE_CRC32 = 0x0000020b         /** 基带热升级校验命令 */
const BB_SET_MCS_MODE = 0x0000020c                  /** 设置MCS控制模式手动、自动 */
const BB_SET_MCS = 0x0000020d                       /** 设置MCS挡位，仅在手动模式下支持 */
const BB_SET_SYS_REBOOT = 0x0000020e                /** 系统重启 */
const BB_SET_MASTER_DEV = 0x0000020f                /** 设置导演模式下的主DEV设备，仅导演模式AP侧支持 */
const BB_SET_FRAME_CHANGE = 0x00000210              /** 运行中帧结构改变(仅1V1模式) */
const BB_SET_COMPLIANCE_MODE = 0x00000211           /** 设置频点合规模式 */
const BB_SET_BAND_MODE = 0x00000212                 /** 设置频段切换模式 */
const BB_SET_BAND = 0x00000213                      /** 设置工作频段 */
const BB_FORCE_CLS_SOCKET_ALL = 0x00000214          /** 强制关闭所有socket 可能会造成socket信息不同步 */
const BB_SET_REMOTE = 0x00000215                    /** 设置通讯对端的配置 */
const BB_SET_BANDWIDTH = 0x00000216                 /** 1V1模式下手动改变bandwidth */
const BB_SET_DFS = 0x00000217                       /** 设置DFS检测 */
const BB_SET_RF = 0x00000218                        /** RF收发开关动态控制 */
const BB_SET_POWER_SAVE_MODE = 0x00000219           /** 1V1模式低功耗策略模式 */
const BB_SET_POWER_SAVE = 0x0000021a                /** 1V1模式低功耗手动周期 */
const BB_SET_REG = 0x00000264                       /** 基带寄存器写入命令字，本类型用于调试诊断 */
const BB_SET_CFG = 0x00000265                       /** 基带配置文件写入命令字 */
const BB_RESET_CFG = 0x00000266                     /** reset基带配置文件命令字 */
const BB_SET_PLOT = 0x00000267                      /** 设置基带plot debug参数 */
const BB_SET_DBG_MODE = 0x00000268                  /** 设置基带SDK进入debug mode（软件不工作） */
const BB_SET_FREQ = 0x00000269                      /** 设置物理用户的发送或接收频率 */
const BB_SET_TX_MCS = 0x0000026a                    /** 设置物理用户的发送MCS */
const BB_SET_RESET = 0x0000026b                     /** 基带RESET命令 */
const BB_SET_TX_PATH = 0x0000026c                   /** 设置无线发射通道，用于功率测试 */
const BB_SET_RX_PATH = 0x0000026d                   /** 设置无线接收通道，用于灵敏度测试 */
const BB_SET_POWER_OFFSET = 0x0000026e              /** 设置功率补偿，用于功率测试 */
const BB_SET_POWER_TEST_MODE = 0x0000026f           /** 进入产测功率测试模式 */
const BB_SET_SENSE_TEST_MODE = 0x00000270           /** 进入产测灵敏度测试模式 */
const BB_SET_ORIG_CFG = 0x00000271                  /** 加载原始镜像配置，仅在基带IDLE状态支持 */
const BB_SET_SINGLE_TONE = 0x00000272               /** 单音信号，需要先进入debug模式 */
const BB_SET_PURE_SLOT = 0x00000273                 /** 纯图传模式，注意暂不支持退出 */
const BB_SET_FACTORY_POWER_OFFSET_SAVE = 0x00000274 /** 产测功率校准保存指令 */
const BB_SET_PRJ_DISPATCH = 0x000002c8              /** 二级SET命令分发 */

// special control
const BB_START_REQ = 0x00000300
const BB_STOP_REQ = 0x00000301
const BB_INIT_REQ = 0x00000302
const BB_DEINIT_REQ = 0x00000303
const BB_RPC_GET_LIST = 0x00000400          /** 获取8030可用列表 */
const BB_RPC_SEL_ID = 0x00000401            /** 选择指定的8030进行通讯 */
const BB_RPC_GET_MAC = 0x00000402           /** 获取指定8030的mac */
const BB_RPC_GET_HOTPLUG_EVENT = 0x00000403 /** 获取设备上下线通知 */
const BB_RPC_SOCK_BUF_STA = 0x00000404      /** 查询socket buff状态 */
const BB_RPC_TEST = 0x00000405              /** 测试服务器连通性 */
const BB_RPC_SERIAL_LIST = 0x00000500       /** 获取串口列表 */
const BB_RPC_SERIAL_SETUP = 0x00000501      /** 设置串口 */
