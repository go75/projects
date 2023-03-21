<script setup lang="ts">
  import {defineComponent, ref} from 'vue'
  // 请求id
  enum ID {
    // 获取好友会话
    GetFriendSession = 0,
    // 获取群聊会话
    GetGroupSession = 1,
    // 获取新好友信息
    GetNewFriend = 2,
    // 获取新群友消息
    GetNewGroup = 3,
    // 获取好友列表
    GetFriendList = 4,
    // 获取群聊列表
    GetGroupList = 5,
    // 添加好友
    AddFriend = 6,
    // 添加群聊
    AddGroup = 7,
    // 发送好友消息
    SendFriendMsg = 8,
    // 发送群聊消息
    SendGroupMsg = 9,
    // 通过用户名称模糊查询用户
    GetFuzzyUserByUserName = 10,
    // 通过群聊名称模糊查询群聊
    GetFuzzyGroupByGroupName = 11,
    // 同意新好友请求
    AgreeNewFriend = 12,
    // 同意新群友请求
    AgreeNewGroup = 13,
    // 拒绝新好友请求
    RefuseNewFriend = 14,
    // 拒绝新群友请求
    RefuseNewGroup = 15,
  }
  // 请求负载类型
  enum Type {
    // 文本
    Text = 0,
    // 图片
    Pic = 1,
    // 音频
    Video = 2,
    // 文件
    File = 3,
  }
  // 请求
  class Request {  
    // 消息id
    ID: ID
    // 负载类型
    Type: Type
    // 接收方id
    ProcessId:number
    // 请求负载
    Payload:string
    constructor(id:ID, type:Type, processId:number, payload:string) {
      this.ID = id
      this.Type = type
      this.ProcessId = processId
      this.Payload = payload
    }
    toJson(): string {
      let s = JSON.stringify(this)
      return s
    }
  }
  // 用户
  class User {
    ID:number
    // 名称
    Name:string
    // 个人介绍
    Introduce:string
    // 构造函数 
    constructor(ID:number,Name:string,Introduce:string) { 
      this.ID = ID
      this.Name = Name
      this.Introduce =Introduce
    }
  }
  // 群聊
  class Group {
    // 群聊id
    ID:number
	  // 群聊名称
	  Name:string
	  // 群主id
	  MasterId:number
	  // 群聊简介
	  Introduce:string
    // 构造函数
    constructor(id:number, name:string, masterId:number, introduce:string) {
      this.ID = id
      this.Name = name
      this.MasterId = masterId
      this.Introduce = introduce
    }
  }
  // 用户聊天记录
  class UserMsg {
    ID:number
    // 发送方id
    SenderId:number
    // 接收方id
    ReceiverId:number
    // 消息类型 文字为0, 图片为1, 音频为2
    Type:number
    // 消息内容
    Content:string
    constructor(id:number, senderId:number, receiverId:number,type:number,content:string) {
      this.ID = id
      this.SenderId = senderId
      this.ReceiverId = receiverId
      this.Type = type
      this.Content = content
    }
  }
  // 聊天消息
  class Msg {
    // 发送方名称
    SenderName:string
    // 消息类型
    Option: Type
    // 消息内容
    Data:string
    constructor(senderName:string, option:number, data:string) {
      this.SenderName = senderName
      this.Option = option
      this.Data = data
    }
  }
  const nav = [{
    url: "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48c3ZnIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Ik00IDZINDRWMzZIMjlMMjQgNDFMMTkgMzZINFY2WiIgZmlsbD0ibm9uZSIgc3Ryb2tlPSIjMzMzIiBzdHJva2Utd2lkdGg9IjQiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjxwYXRoIGQ9Ik0yMyAyMUgyNS4wMDI1IiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIi8+PHBhdGggZD0iTTMzLjAwMSAyMUgzNC45OTk5IiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIi8+PHBhdGggZD0iTTEzLjAwMSAyMUgxNC45OTk5IiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIi8+PC9zdmc+",
    fn: function() {
      index.value = 1
    }
  },{
    url: "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48c3ZnIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Ik04IDM2VjQ0SDQwVjRIOFYxMiIgc3Ryb2tlPSIjMzMzIiBzdHJva2Utd2lkdGg9IjQiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjxwYXRoIGQ9Ik02IDMwSDEwIiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIiBzdHJva2UtbGluZWpvaW49InJvdW5kIi8+PHBhdGggZD0iTTYgMjRIMTAiIHN0cm9rZT0iIzMzMyIgc3Ryb2tlLXdpZHRoPSI0IiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiLz48cGF0aCBkPSJNNiAxOEgxMCIgc3Ryb2tlPSIjMzMzIiBzdHJva2Utd2lkdGg9IjQiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjxjaXJjbGUgY3g9IjI0IiBjeT0iMTciIHI9IjQiIGZpbGw9Im5vbmUiIHN0cm9rZT0iIzMzMyIgc3Ryb2tlLXdpZHRoPSI0IiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiLz48cGF0aCBkPSJNMzIgMzVDMzIgMzAuNTgxNyAyOC40MTgzIDI3IDI0IDI3QzE5LjU4MTcgMjcgMTYgMzAuNTgxNyAxNiAzNSIgc3Ryb2tlPSIjMzMzIiBzdHJva2Utd2lkdGg9IjQiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjwvc3ZnPg==",
    fn: function() {
      index.value = 2
    }
  },{
    url: "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48c3ZnIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Ik0xOSAyMEMyMi44NjYgMjAgMjYgMTYuODY2IDI2IDEzQzI2IDkuMTM0MDEgMjIuODY2IDYgMTkgNkMxNS4xMzQgNiAxMiA5LjEzNDAxIDEyIDEzQzEyIDE2Ljg2NiAxNS4xMzQgMjAgMTkgMjBaIiBmaWxsPSJub25lIiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNMzYgMjlWNDFWMjlaIiBmaWxsPSJub25lIi8+PHBhdGggZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiIGQ9Ik0zMCAzNUg0MkgzMFoiIGZpbGw9Im5vbmUiLz48cGF0aCBkPSJNMzYgMjlWNDFNMzAgMzVINDIiIHN0cm9rZT0iIzMzMyIgc3Ryb2tlLXdpZHRoPSI0IiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiLz48cGF0aCBkPSJNMjcgMjhIMTguOEMxNC4zMTk2IDI4IDEyLjA3OTQgMjggMTAuMzY4MSAyOC44NzE5QzguODYyNzggMjkuNjM4OSA3LjYzODkzIDMwLjg2MjggNi44NzE5NSAzMi4zNjgxQzYgMzQuMDc5NCA2IDM2LjMxOTYgNiA0MC44VjQySDI3IiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIiBzdHJva2UtbGluZWpvaW49InJvdW5kIi8+PC9zdmc+",
    fn: function() {
      index.value = 3
    }
  },{
    url: "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48c3ZnIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Ik0xOC4yODM4IDQzLjE3MTNDMTQuOTMyNyA0Mi4xNzM2IDExLjk0OTggNDAuMzIxMyA5LjU4Nzg3IDM3Ljg2N0MxMC40NjkgMzYuODIyNyAxMSAzNS40NzM0IDExIDM0LjAwMDFDMTEgMzAuNjg2NCA4LjMxMzcxIDI4LjAwMDEgNSAyOC4wMDAxQzQuNzk5NTUgMjguMDAwMSA0LjYwMTM5IDI4LjAxIDQuNDA1OTkgMjguMDI5MkM0LjEzOTc5IDI2LjcyNzcgNCAyNS4zODAzIDQgMjQuMDAwMUM0IDIxLjkwOTUgNC4zMjA3NyAxOS44OTM4IDQuOTE1NzkgMTcuOTk5NUM0Ljk0MzgxIDE3Ljk5OTkgNC45NzE4OCAxOC4wMDAxIDUgMTguMDAwMUM4LjMxMzcxIDE4LjAwMDEgMTEgMTUuMzEzOCAxMSAxMi4wMDAxQzExIDExLjA0ODggMTAuNzc4NiAxMC4xNDkzIDEwLjM4NDYgOS4zNTAxMUMxMi42OTc1IDcuMTk5NSAxNS41MjA1IDUuNTkwMDIgMTguNjUyMSA0LjcyMzE0QzE5LjY0NDQgNi42NjgxOSAyMS42NjY3IDguMDAwMTMgMjQgOC4wMDAxM0MyNi4zMzMzIDguMDAwMTMgMjguMzU1NiA2LjY2ODE5IDI5LjM0NzkgNC43MjMxNEMzMi40Nzk1IDUuNTkwMDIgMzUuMzAyNSA3LjE5OTUgMzcuNjE1NCA5LjM1MDExQzM3LjIyMTQgMTAuMTQ5MyAzNyAxMS4wNDg4IDM3IDEyLjAwMDFDMzcgMTUuMzEzOCAzOS42ODYzIDE4LjAwMDEgNDMgMTguMDAwMUM0My4wMjgxIDE4LjAwMDEgNDMuMDU2MiAxNy45OTk5IDQzLjA4NDIgMTcuOTk5NUM0My42NzkyIDE5Ljg5MzggNDQgMjEuOTA5NSA0NCAyNC4wMDAxQzQ0IDI1LjM4MDMgNDMuODYwMiAyNi43Mjc3IDQzLjU5NCAyOC4wMjkyQzQzLjM5ODYgMjguMDEgNDMuMjAwNSAyOC4wMDAxIDQzIDI4LjAwMDFDMzkuNjg2MyAyOC4wMDAxIDM3IDMwLjY4NjQgMzcgMzQuMDAwMUMzNyAzNS40NzM0IDM3LjUzMSAzNi44MjI3IDM4LjQxMjEgMzcuODY3QzM2LjA1MDIgNDAuMzIxMyAzMy4wNjczIDQyLjE3MzYgMjkuNzE2MiA0My4xNzEzQzI4Ljk0MjggNDAuNzUyIDI2LjY3NiAzOS4wMDAxIDI0IDM5LjAwMDFDMjEuMzI0IDM5LjAwMDEgMTkuMDU3MiA0MC43NTIgMTguMjgzOCA0My4xNzEzWiIgZmlsbD0ibm9uZSIgc3Ryb2tlPSIjMzMzIiBzdHJva2Utd2lkdGg9IjQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiLz48cGF0aCBkPSJNMjQgMzFDMjcuODY2IDMxIDMxIDI3Ljg2NiAzMSAyNEMzMSAyMC4xMzQgMjcuODY2IDE3IDI0IDE3QzIwLjEzNCAxNyAxNyAyMC4xMzQgMTcgMjRDMTcgMjcuODY2IDIwLjEzNCAzMSAyNCAzMVoiIGZpbGw9Im5vbmUiIHN0cm9rZT0iIzMzMyIgc3Ryb2tlLXdpZHRoPSI0IiBzdHJva2UtbGluZWpvaW49InJvdW5kIi8+PC9zdmc+",
    fn: function() {
      index.value = 4
    }
  }]
 
  // 好友列表
  let friendListValue:User[] = []
  var friendList = ref(friendListValue)
  // 群聊列表
  let groupListValue:Group[] = []
  var groupList = ref(groupListValue)
  // 用户消息列表
  let userMsgListValue:UserMsg[] = []
  var userMsgList = ref(userMsgListValue)
  // 消息列表
  let msgListValue:Msg[] = []
  var msgList = ref(msgListValue)
  // 新好友列表
  let newFriendListValue:User[] = []
  var newFriendList = ref(newFriendListValue)
  // 新群友列表
  let newGroupListValue:Group[] = []
  var newGroupList = ref(newGroupListValue)
  // 搜索用户列表
  let searchUserListValue:User[] = []
  var searchUserList = ref(searchUserListValue)
  // 搜索群聊列表
  let searchGroupListValue:Group[] = []
  var searchGroupList = ref(searchGroupListValue)
  var index = ref(0)
  let userValue:User = new User(0, '', '')
  var otherUser = ref(userValue)
  var searchInput:string
  var msgInput:string
  var pc:RTCPeerConnection
  const myid = parseInt(localStorage.getItem('id') as string)
  const myname = localStorage.getItem('name') as string
  const v1 = ref()
  // websocket模块 ================================================================================================================
  const ws:WebSocket=new WebSocket("ws://127.0.0.1:9999/upgrade")
  ws.onopen = function (evt) {
      alert('连接建立')
      // 获取好友列表
      send(ID.GetFriendList, Type.Text, 0, '')
      // 获取群聊列表
      send(ID.GetGroupList, Type.Text, 0, '')
      // 获取新好友列表
      send(ID.GetNewFriend, Type.Text, 0, '')
      // 获取新群友列表
      send(ID.GetNewGroup, Type.Text, 0, '')
  }
  ws.onclose = function (evt) {
      alert('连接断开')
  }
  ws.onmessage = function (evt) {
      let req:Request = JSON.parse(evt.data)
      switch (req.ID) {
        case ID.GetFriendSession:
          // 获取好友会话
          userMsgList.value = JSON.parse(req.Payload)
          console.log(userMsgList)
          break
        case ID.GetGroupSession:
          // 获取群聊会话
          break
        case ID.GetNewFriend:
          // 获取新好友列表
          newFriendList.value = JSON.parse(req.Payload)
          console.log('new friend list: ', JSON.parse(req.Payload))
          break
        case ID.GetNewGroup:
          // 获取新群友列表
          newGroupList.value = JSON.parse(req.Payload)
          break
        case ID.GetFriendList:
          // 更新好友列表
          
          friendList.value = JSON.parse(req.Payload)
          break
        case ID.GetGroupList:
          // 更新群聊列表
          groupList.value = JSON.parse(req.Payload)
          break
        case ID.AddFriend:
          // 添加好友
          break
        case ID.AddGroup:
          // 添加群聊
          break
        case ID.SendFriendMsg:
          // 发送好友消息

          break
        case ID.SendGroupMsg:
          // 发送群聊消息
          new RTCPeerConnection()
          break
        case ID.GetFuzzyUserByUserName:
          searchUserList.value = JSON.parse(req.Payload)
          break
        case ID.GetFuzzyGroupByGroupName:
          searchGroupList.value = JSON.parse(req.Payload)
          break
      }
  }
  ws.onerror = function (evt) {
      alert("ERROR")
  }
  function send(id:ID, type:Type, processId:number, payload:string) {
    let data = new Request(id, type, processId, payload).toJson()
    console.log(data)
    ws.send(data)
  }
  // ============================================================================================================================
  // 获取用户会话
  function sessionFn(friend:User) {
    send(ID.GetFriendSession, Type.Text, friend.ID, '')
    otherUser.value = friend
  }

  function searchFn(name:string) {
    send(ID.GetFuzzyUserByUserName, Type.Text, 0, name)
  }

  function addUserFn(id:number) {
    send(ID.AddFriend, Type.Text, id, '')
  }

  function agreeNewFriend(id:number, i:number) {
    send(ID.AgreeNewFriend, Type.Text, id, '')
    friendList.value.push(newFriendList.value[i])
    delete newFriendList.value[i]
  }

  function refuseNewFriend(id:number, i:number) {
    send(ID.RefuseNewFriend, Type.Text, id, '')
    delete newFriendList.value[i]
  }

  function agreeNewGroup(id:number, i:number) {
    send(ID.AgreeNewGroup, Type.Text, id, '')
    groupList.value.push(newGroupList.value[i])
    delete newGroupList.value[i]
  }

  function refuseNewGroup(id:number, i:number) {
    send(ID.RefuseNewGroup, Type.Text, id, '')
    delete newGroupList.value[i]
  }
  var constraints = {
    audio: true,
    video: true,
  }
  function openCamera() {
      navigator.mediaDevices.getUserMedia(constraints).then(function(stream){
        v1.value.srcObject = stream
      }).catch(function(err){
        console.error("get user media err: " + err)
      })
  }
  function closeCamera() {
    var tracks = v1.value.srcObject.getTracks()
    for(var i = 0 ; i< tracks.length ; i++) {
        tracks[i].stop()
    }
  }
</script>

<template>
  <div class="home">
    <div class="main">
      <div class="nav">
        <img class="head-icon" :src="'http://127.0.0.1:9999/head/'+myname+'.png'">
        <img class="head-icon" :src="img.url" @click="img.fn()" v-for="img in nav">
      </div>
      <div class="content">
        <div class="userlist">

          <div class="userlist-left">
            <div class="msg" v-for="friend in friendList" @click="sessionFn(friend)">
              <img class="msg-img" :src="'http://127.0.0.1:9999/head/'+friend.Name+'.png'">
              <p class="msg-title">
                {{ friend.Name }}
              </p>
            </div>
          </div>

          <div class="userlist-right" v-if="index==1">
            <div class="chat-title">
              <h1>{{ otherUser.Name }}</h1>
            </div>
            <div class="chats">
              <div :class="msg.SenderId==myid?'chat-right':'chat-left'" v-for="msg in userMsgList">
                <img :src="'http://127.0.0.1:9999/head/' + (msg.SenderId==myid?myname:otherUser.Name) +'.png'" class="list-img">
                <!-- 文本消息 -->
                <h3 v-if="msg.Type==Type.Text" :class="msg.SenderId==myid?'right-msg':'left-msg'">
                  {{ msg.Content }}
                </h3>
                <!-- 图片消息 -->
                <img v-else-if="msg.Type==Type.Pic" :src="msg.Content" :class="msg.SenderId==myid?'right-msg':'left-msg'">
              </div>
            </div>
            <div class="chat-boom">
              <input type="text" v-model="msgInput">
              <!-- 视频 -->
              <img class="input-btn" @click="" src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48c3ZnIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Ik0xNC4zNzU2IDkuNzk0MUMxNS4xMDIyIDkuNzk0MSAxNS43NzE2IDEwLjE4ODEgMTYuMTI0MiAxMC44MjM0TDE4LjU3MDYgMTUuMjMwMkMxOC44OTEgMTUuODA3MiAxOC45MDYgMTYuNTA1MSAxOC42MTA5IDE3LjA5NTRMMTYuMjU0IDIxLjgwOUMxNi4yNTQgMjEuODA5IDE2LjkzNyAyNS4zMjA0IDE5Ljc5NTUgMjguMTc4OUMyMi42NTM5IDMxLjAzNzMgMjYuMTUzNSAzMS43MDg1IDI2LjE1MzUgMzEuNzA4NUwzMC44NjY1IDI5LjM1MjFDMzEuNDU3MiAyOS4wNTY3IDMyLjE1NTUgMjkuMDcyIDMyLjczMjcgMjkuMzkyOUwzNy4xNTIxIDMxLjg0OTlDMzcuNzg2NyAzMi4yMDI3IDM4LjE4MDIgMzIuODcxOCAzOC4xODAyIDMzLjU5NzlMMzguMTgwMiAzOC42NzE0QzM4LjE4MDIgNDEuMjU1MSAzNS43ODAzIDQzLjEyMTEgMzMuMzMyMyA0Mi4yOTUxQzI4LjMwNDMgNDAuNTk4NiAyMC40OTk3IDM3LjM2ODQgMTUuNTUyOCAzMi40MjE1QzEwLjYwNiAyNy40NzQ3IDcuMzc1NzYgMTkuNjcgNS42NzkyMiAxNC42NDIxQzQuODUzMiAxMi4xOTQgNi43MTkyOSA5Ljc5NDEgOS4zMDI5MyA5Ljc5NDFMMTQuMzc1NiA5Ljc5NDFaIiBmaWxsPSJub25lIiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjxwYXRoIGQ9Ik0zOSAxNUgyN1Y1SDM5VjhMNDQgNlYxNEwzOSAxMlYxNVoiIGZpbGw9Im5vbmUiIHN0cm9rZT0iIzMzMyIgc3Ryb2tlLXdpZHRoPSI0IiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiLz48L3N2Zz4=">
              <!-- 表情 -->
              <img class="input-btn" @click="" src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48c3ZnIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Ik0yNCA0NEMzNS4wNDU3IDQ0IDQ0IDM1LjA0NTcgNDQgMjRDNDQgMTIuOTU0MyAzNS4wNDU3IDQgMjQgNEMxMi45NTQzIDQgNCAxMi45NTQzIDQgMjRDNCAzNS4wNDU3IDEyLjk1NDMgNDQgMjQgNDRaIiBmaWxsPSJub25lIiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjxwYXRoIGQ9Ik0zMSAxOFYxOSIgc3Ryb2tlPSIjMzMzIiBzdHJva2Utd2lkdGg9IjQiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjxwYXRoIGQ9Ik0xNyAxOFYxOSIgc3Ryb2tlPSIjMzMzIiBzdHJva2Utd2lkdGg9IjQiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjxwYXRoIGQ9Ik0zMSAzMUMzMSAzMSAyOSAzNSAyNCAzNUMxOSAzNSAxNyAzMSAxNyAzMSIgc3Ryb2tlPSIjMzMzIiBzdHJva2Utd2lkdGg9IjQiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjwvc3ZnPg==">
              <!-- 语音 -->
              <img class="input-btn" @click="" src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48c3ZnIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Ik0yNCA0NEMzNS4wNDU3IDQ0IDQ0IDM1LjA0NTcgNDQgMjRDNDQgMTIuOTU0MyAzNS4wNDU3IDQgMjQgNEMxMi45NTQzIDQgNCAxMi45NTQzIDQgMjRDNCAzNS4wNDU3IDEyLjk1NDMgNDQgMjQgNDRaIiBmaWxsPSJub25lIiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjxwYXRoIGQ9Ik0xNyAyNS44OTk0QzE4LjEwNDYgMjUuODk5NCAxOSAyNS4wMDQgMTkgMjMuODk5NEMxOSAyMi43OTQ4IDE4LjEwNDYgMjEuODk5NCAxNyAyMS44OTk0QzE1Ljg5NTQgMjEuODk5NCAxNSAyMi43OTQ4IDE1IDIzLjg5OTRDMTUgMjUuMDA0IDE1Ljg5NTQgMjUuODk5NCAxNyAyNS44OTk0WiIgZmlsbD0iIzMzMyIvPjxwYXRoIGQ9Ik0yMS45NDk3IDI4Ljg0OTJDMjMuMjE2NSAyNy41ODI1IDI0IDI1LjgzMjUgMjQgMjMuODk5NUMyNCAyMS45NjY1IDIzLjIxNjUgMjAuMjE2NSAyMS45NDk3IDE4Ljk0OTciIHN0cm9rZT0iIzMzMyIgc3Ryb2tlLXdpZHRoPSI0IiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiLz48cGF0aCBkPSJNMjYuODk5NCAzMy43OTlDMjkuNDMyOSAzMS4yNjU1IDMwLjk5OTkgMjcuNzY1NSAzMC45OTk5IDIzLjg5OTVDMzAuOTk5OSAyMC4wMzM1IDI5LjQzMjkgMTYuNTMzNSAyNi44OTk0IDE0IiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIiBzdHJva2UtbGluZWpvaW49InJvdW5kIi8+PC9zdmc+">
              <!-- 文件 -->
              <img class="input-btn" @click="" src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48c3ZnIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Ik00MCAyM1YxNEwzMSA0SDEwQzguODk1NDMgNCA4IDQuODk1NDMgOCA2VjQyQzggNDMuMTA0NiA4Ljg5NTQzIDQ0IDEwIDQ0SDIyIiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIiBzdHJva2UtbGluZWpvaW49InJvdW5kIi8+PHBhdGggZD0iTTMzIDI5VjQzIiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIiBzdHJva2UtbGluZWpvaW49InJvdW5kIi8+PHBhdGggZD0iTTI2IDM2SDMzSDQwIiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIiBzdHJva2UtbGluZWpvaW49InJvdW5kIi8+PHBhdGggZD0iTTMwIDRWMTRINDAiIHN0cm9rZT0iIzMzMyIgc3Ryb2tlLXdpZHRoPSI0IiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiLz48L3N2Zz4=">
              <!-- 发送 -->
              <img class="input-btn" @click="function(){
                send(ID.SendFriendMsg, Type.Text, otherUser.ID, msgInput)
                userMsgList.push(new UserMsg(0, myid, otherUser.ID, Type.Text, msgInput))
                msgInput = ''
                }" src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48c3ZnIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Ik00MiA2TDQgMjAuMTM4M0wyNCAyNC4wMDgzTDI5LjAwNTIgNDRMNDIgNloiIHN0cm9rZT0iIzMzMyIgc3Ryb2tlLXdpZHRoPSI0IiBzdHJva2UtbGluZWpvaW49InJvdW5kIi8+PHBhdGggZD0iTTI0LjAwODMgMjQuMDA4NEwyOS42NjUxIDE4LjM1MTYiIHN0cm9rZT0iIzMzMyIgc3Ryb2tlLXdpZHRoPSI0IiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiLz48L3N2Zz4=">
            </div>
          </div>

          <div class="userlist-right" v-else-if="index==2">
            <h1>视频通话</h1>
            <video ref="v1" autoplay playsinline muted controls></video>
            <button @click="openCamera()">打开摄像头</button>
            <button @click="closeCamera()">关闭摄像头</button>
          </div>

          <div class="userlist-right" v-else-if="index==3">
            <input type="text" v-model="searchInput">
            <button @click="searchFn(searchInput)">搜索</button>
            <div class="new-list">
              <div v-for="user in searchUserList">
                <button @click="addUserFn(user.ID)">添加</button>
                <img class="msg-img" :src="'http://127.0.0.1:9999/head/'+user.Name+'.png'">
                {{ user.Name }}
              </div>
            </div>
            <div class="new-list">
              <h1>新朋友</h1>
              <div v-for="user, i in newFriendList">
                <button @click="agreeNewFriend(user.ID, i)">同意</button>
                <button @click="refuseNewFriend(user.ID, i)">拒绝</button>
                <img class="msg-img" :src="'http://127.0.0.1:9999/head/'+user.Name+'.png'">
                {{ user.Name }}
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>