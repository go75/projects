<template>
  <div class="content">
    <h1>速传</h1>
    <br>
    <select id="options">
      <option v-for="address in addresses" :value="address">
        {{ address }}
      </option>
    </select>
    <br>
    <input id="upload" type="file" :name="filenamae">
    <br>
    <button @click="upload">上传</button>
    <img :src="url">
  </div>
</template>

<script setup lang="ts">
  import { ref } from 'vue'
  const xhr = new XMLHttpRequest()
  var filenamae = ref("")
  var url  = ref("")
  var addresses = ref([""])
  function upload() {
    doReq('post', 'http://127.0.0.1:8888/file/1.html', function(text:string){
      url.value = text
      console.log(url.value)
    }, function(text:string){})
    // let file = document.getElementsByTagName('input')[0].files?.item(0)
    // // 创建读取器
    // let reader=new FileReader()
    // // 开始读取
    // reader.readAsText(file as File)
    // let path = file?.name as string
    // // 监听文件的读取状态
    // reader.onload=function(){
    // let ss = reader.result as string
    // console.log(ss)
    // let s = window.atob(ss)
    // url.value = s
    // path += "/" + reader.result as string
    // console.log(path)
      // doReq('post', 'http://127.0.0.1:8888/file/'+file?.name, function(text:string) {
      //   console.log(text)
      // }, function (text:string) {
      //   alert(text)
      // })
    //}
  }

  function doReq(method:string, url:string, ok:(text:string)=>void, err:(text:string)=>void) {
      xhr.open(method, url, false)
      xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded')
      xhr.send("")
      if (xhr.status == 200) {
          ok(xhr.responseText)
      } else {
          err(xhr.responseText)
      }
  }
  // window.onload = function() {
  //   doReq('get', 'http://127.0.0.1:8888/addresses', function(text:string){
  //     addresses.value = JSON.parse(text).addresses
  //   }, function(text:string){
  //     alert(text)
  //   })
  // }

  const getFilePath:(file:File)=>string = (file:File) => {
    let path = ""
    if (window.URL != undefined) {
        // mozilla(firefox)
        path = window.URL.createObjectURL(file);
    } else if (window.webkitURL != undefined) {
        // webkit or chrome
        path = window.webkitURL.createObjectURL(file);
    }
    return path
  }
</script>

<style>
  .content {
    text-align: center;
  }
</style>