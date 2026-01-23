
帮我生成PPT：

1，今年主要做了低代码的运营者 管理员模块（你可以理解成后台管理模块） 低代码的数据关联部分模块
2，主要合并了赣政通pc端，开发了赣政通的锁屏模块，云控，人保的linux的代理，组织结构的左侧树结构改造
3，处理加班类型的需求，完成工作圈的黑名单开发，修复考勤加班等bug
4，开发关保平台的权限后台，移动端的列表显示
5，开发水利部访客系统的访客系统
AI结合

AI 工作汇报

1. AI 的着力点在于“把 0% 的想法快速推进到 30% 可运行原型”。它擅长“已有范式内的局部加速”，不擅长“从无到有的范式突破”。因此，我们要用高质量 Prompt 把“需求边界、数据格式、交互范式”一次性描述清楚，让模型在 30% 框架内填代码，而非凭空创新。

   示范：生成一个管理后台页面的 Prompt（可直接投喂给模型）

   1.1 顶部导航栏

   - 用 el-menu，背景色 #409EFF，文字白色，选中时背景加深 10%。
   - 左侧预留项目名称插槽（slot="brand"）。
   - 中间菜单项通过数组 NavItems 注入，结构：
     ```ts
     interface NavItem { name: string; route: string; icon?: string }
     ```
   - 暴露事件：@select、@menu-click。
   - Store 的 module 划分、state 形状、mutation 名由人工先锁定，AI 只生成具体 CRUD 逻辑。

   1.2 左侧列表区块

   - 点击高亮用 class 切换（.is-active）。
   - 中间内容区 flex:1 自适应。

   1.3 左上角按钮组

   - 添加按钮：el-button，左侧内置 Plus 图标，文字“新增”。
   - 搜索框：el-input，左侧内置 Search 图标，v-model 绑定 keyword，触发 @search。

   1.4 列表实现

   - 字段校对由人工完成；让 AI 直接生成固定列数的 el-table，方便后续填充。
   - 要求：给出 `<el-table-column>` 骨架，prop 与后端字段保持一致。

   1.5 封装 ListTable 组件

   - 外部传入 columns: `{ key, label, width?, slot? }[]`。
   - 内部维护 rendererMap: `{ [key:string]: Function }`，按 column.key 直接映射渲染函数，降低 if-else。
   - 插槽命名规则：`table-[key]`，例如 `<template #table-status>`。
2. 用 AI 提升学习能力把原来“查文档 + 写样板”的时间省下来，让 AI 在 30 分钟内给出“带设计模式 + 注释 + 单测模板”的高可维护代码。做法：

   - 先指定设计模式（如 Strategy、Factory、Observer）。
   - 要求 AI 生成 TypeScript 实现，并逐行补充 JSDoc 与单测用例。
   - 对存量代码，可让 AI 按相同模式做“渐进式重构”，人工只做 Code Review。
3. LangChain 调研结论

   3.1 Function Calling

   - 本地预注册函数，LLM 仅返回 JSON 参数；本地框架负责反射调用。
   - LLM 本身不拥有、也不执行函数——它只是“点名 + 传参”。

   3.2 Multi-Agent 协作

   - 同一进程内可启多个 Agent 实例，每个实例 system prompt 只聚焦单一职责（产品、开发、测试）。
   - 通过消息队列顺序传递上下文，降低单 Agent 的上下文长度与幻觉概率。

   3.3 MCP（Model Context Protocol）

   - MCP 是 OpenAI 提出的统一“工具描述 + 调用”协议，Agent 通过 MCP 把工具清单发给 LLM。
   - LLM 按 MCP 格式返回调用指令，本地 SDK 再执行真正的 Function Calling。
   - 类比：Agent 是顾客，MCP 是菜单 + 点单协议，Function Calling 是后厨实际炒菜。
     生成有求 我只是一个开发

1，工作内容 简介干练，突出我做了什么即可，不需要图片进行填充 最好是分布展示
2，ai结合部分，需要图片 让人知道言之何物
