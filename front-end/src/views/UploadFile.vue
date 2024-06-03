<template>
  <div style="height: 100%" class="body-container">
    <!-- Header -->
    <div class="upload-header">
      <Header />
    </div>
    <!-- 文件上传和文本输入按钮 -->
    <div class="upload-choice-option">
      <span class="choice-span span1" @click="ChangeType(0)">文件上传</span>
      <span class="choice-span span2" @click="ChangeType(1)">文本输入</span>
    </div>
    <!-- 上传内容 -->
    <div class="upload-file-or-text">
      <!-- 文件上传 -->
      <div class="uploadFile" v-if="typeView === 0">
        <dv-border-box-11 title="文件上传">
          <Upload
            class="upload-file"
            multiple
            type="drag"
            :before-upload="beforeUpload"
            action=""
            :show-file-list="false"
          >
            <div class="upload-file-inner" style="padding: 20px 0">
              <Icon
                type="ios-cloud-upload"
                size="52"
                style="color: #3399ff"
              ></Icon>
              <p>上传txt文件</p>
            </div>
          </Upload>
        </dv-border-box-11>
      </div>
      <!-- 文本输入 -->
      <div class="uploadText" v-if="typeView === 1">
        <dv-border-box-11 title="文本输入">
          <Input
            class="upload-text"
            type="textarea"
            :autosize="{ maxRows: 10, minRows: 10 }"
            :rows="10"
            placeholder="请输入您需要鉴别的新闻"
          />
        </dv-border-box-11>
      </div>
    </div>
    <!-- 提交按钮 -->
    <div class="upload-confirm" @click="RouteToAnalysisResult">
      <button class="send-btn">确认上传</button>
    </div>
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
export default {
  name: 'uploadFile',
  components: {
    Header
  },
  data() {
    return {
      typeView: 0
    }
  },
  methods: {
    ChangeType(type) {
      this.typeView = type
    },
    RouteToAnalysisResult() {
      this.$router.push('/analysisResult')
    },
    ReadTxtContent() {
      let xhr = new XMLHttpRequest()
      let okStatus = document.location.protocol === 'file:' ? 0 : 200
      xhr.open('GET', 'text.txt', false) // public文件夹下的绝对路径
      xhr.overrideMimeType('text/html;charset=utf-8')
      xhr.send(null)
      console.log(xhr.responseText) // xhr.responseText为文本中的内容
    },
    beforeUpload(file) {
      console.log(file)
      const isExe = file.name.endsWith('.txt')
      if (!isExe) {
        this.$message.error('只能上传txt文件')
        return false
      }
      // this.handleChange(file)
      return false
    },
    handleChange(file) {
      let reader = new FileReader() //先new 一个读文件的对象 FileReader
      console.log('读取文件')
      //  reader.readAsText(file.raw, "gb2312");  //读.txt文件
      reader.readAsArrayBuffer(file.raw) //读任意文件
      reader.onload = function (e) {
        var ints = new Uint8Array(e.target.result) //要使用读取的内容，所以将读取内容转化成Uint8Array
        ints = ints.slice(0, 5000) //截取一段读取的内容
        let snippets = new TextDecoder('gb2312').decode(ints) //二进制缓存区内容转化成中文（即也就是读取到的内容）
        console.log('读取的内容如下：')
        console.log(snippets)
      }
    }
  }
}
</script>

<style lang="less">
.body-container {
  background: url('../assets/footer-animation-bg.svg') center center;
  background-color: #050828;
}
.upload-header {
  height: 106px;
  justify-content: space-between;
  align-items: center;
}

.upload-choice-option {
  margin-top: 100px;
  height: 130px;
  text-align: center;

  .choice-span {
    width: 400px;
    border: 0;
    margin: 100px;
    text-transform: uppercase;
    font-size: 20px;
    font-weight: bold;
    padding: 15px 50px;
    border-radius: 50px;
    color: white;
    outline: none;
    position: relative;
  }

  .choice-span:before {
    content: '';
    display: block;
    background: linear-gradient(
      to left,
      rgba(255, 255, 255, 0) 50%,
      rgba(255, 255, 255, 0.4) 50%
    );
    background-size: 210% 100%;
    background-position: right bottom;
    height: 100%;
    width: 100%;
    position: absolute;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    border-radius: 50px;
    transition: all 1s;
    -webkit-transition: all 1s;
  }

  .span1 {
    cursor: pointer;
    background-image: linear-gradient(to right, #25aae1, #40e495);
    box-shadow: 0 7px 15px 0 rgba(49, 196, 190, 0.75);
  }

  .span2 {
    cursor: pointer;
    background-image: linear-gradient(to right, #6253e1, #852d91);
    box-shadow: 0 7px 15px 0 rgba(236, 116, 149, 0.75);
  }

  .span1:hover:before {
    background-position: left bottom;
  }

  .span2:hover:before {
    background-position: left bottom;
  }
}

.upload-file-or-text {
  display: flex;
  justify-content: center;
  align-items: center;

  .uploadFile {
    width: 600px;
    height: 400px;
    position: relative;

    .upload-file {
      padding-top: 160px;
      width: 400px;
      position: absolute;
      top: 35%;
      left: 50%;
      transform: translate(-50%, -50%);

      .ivu-upload-drag {
        border: 0px;
      }

      .ivu-upload:hover {
        border: 0px;
      }

      .upload-file-inner {
        padding: 0px;
        background-color: #050828;

        p {
          font-size: 20px;
          color: white;
        }
      }
    }
  }

  .uploadText {
    width: 600px;
    height: 400px;
    position: relative;

    .upload-text {
      width: 90%;
      position: absolute;
      top: 55%;
      left: 50%;
      transform: translate(-50%, -50%);

      .ivu-input {
        color: skyblue;
        border: 0px;
        background-color: #050d28;
        font-size: 18px;
        text-indent: 2em;
      }

      .ivu-input::placeholder {
        color: skyblue;
      }

      .ivu-input:hover {
        box-shadow: 0 0 0 1px white;
      }

      .ivu-input:focus {
        box-shadow: 0 0 0 1px white;
      }
    }
  }
}

.upload-confirm {
  margin-top: 20px;
  display: flex;
  justify-content: center;
  align-items: center;

  .send-btn {
    display: flex;
    align-items: center;
    justify-content: space-around;
    height: 50px;
    background: #031628;
    border: 1px solid;
    background-color: transparent;
    text-transform: uppercase;
    font-size: 20px;
    padding: 10px 20px;
    font-weight: 300;
    color: #4cc9f0;
    border-radius: 10px;
  }

  .send-btn:hover {
    cursor: pointer;
    color: white;
    border: 0;
    background-color: #4cc9f0;
    -webkit-box-shadow: 10px 10px 99px 6px rgba(76, 201, 240, 1);
    -moz-box-shadow: 10px 10px 99px 6px rgba(76, 201, 240, 1);
    box-shadow: 10px 10px 99px 6px rgba(76, 201, 240, 1);
  }
}
</style>