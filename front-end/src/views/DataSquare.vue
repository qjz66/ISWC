<template>
  <div class="body-container">
    <div class="header">
      <Header />
    </div>
    <!-- 动态部分 -->
    <div class="info-section">
      <!-- 上传动态 -->
      <div class="info-header">
        <el-input
          type="textarea"
          resize="none"
          :rows="2"
          placeholder="请输入需要搜索的关键词"
          v-model="textarea"
        >
        </el-input>
        <div class="btn-control">
          <span class="cancel" @click="cancelSearch">取消</span>
          <el-button class="btn" type="success" round @click="commitSearch()"
            >搜索</el-button
          >
        </div>
      </div>
      <!-- 动态显示 -->
      <div class="info-body" v-for="item in comments" :key="item.date">
        <div class="info-show" v-if="item.showStatus">
          <!-- 个人信息 -->
          <div class="info">
            <span class="info-name">{{ item.fromName }}</span>
            <span class="info-date">{{ item.date }}</span>
          </div>
          <!-- 评论内容 -->
          <div class="content">
            <div class="content-about">{{ item.content }}</div>
          </div>
          <!-- 评论、点赞 -->
          <div class="control">
            <!-- 点赞 -->
            <span class="control-dianzan" v-if="item.likeStatus === 0">
              <img
                src="../picture/dianzan_white.png"
                @click="ChangeLike(item, 1)"
                alt=""
              />
              <span>{{ item.likeNum }}</span>
            </span>
            <span class="control-dianzan" v-else-if="item.likeStatus === 1">
              <img
                src="../picture/dianzan_red.png"
                @click="ChangeLike(item, 0)"
                alt=""
              />
              <span>{{ item.likeNum }}</span>
            </span>
            <!-- 回复 -->
            <span class="control-pinglun" @click="showCommentInput(item)">
              <img src="../picture/pinglun.png" alt="" />
            </span>

            <span
              class="control-agree"
              :style="item.agree == 0 ? 'color : red;' : ''"
              @click="changeAgree(item, 0)"
              >赞成</span
            >
            <span
              class="control-disagree"
              :style="item.agree == 1 ? 'color : red;' : ''"
              @click="changeAgree(item, 1)"
              >反对</span
            >
          </div>
          <!-- 回复列表 -->
          <div class="reply">
            <!-- 评论列表 -->
            <div class="item" v-for="reply in item.reply" :key="reply.date">
              <div class="reply-content">
                <span class="from-name">{{ reply.fromName }}</span
                ><span>: </span>
                <span class="to-name">@{{ reply.toName }} </span>
                <span>{{ reply.content }}</span>
              </div>
              <div class="reply-bottom">
                <span class="reply-bottom-date">{{ reply.date }}</span>
                <span class="reply-text" @click="showCommentInput(item, reply)">
                  <img src="../picture/pinglun.png" alt="" />
                </span>
              </div>
            </div>
            <!-- 评论的输入框 -->
            <transition name="fade">
              <div class="input-wrapper" v-if="showItemId === item.id">
                <el-input
                  v-model="inputComment"
                  type="textarea"
                  resize="none"
                  :rows="3"
                  autofocus
                  placeholder="写下你的评论"
                >
                </el-input>
                <div class="btn-control">
                  <span class="cancel" @click="cancelComment">取消</span>
                  <el-button
                    class="btn"
                    type="success"
                    round
                    @click="commitComment(item)"
                    >确定</el-button
                  >
                </div>
              </div>
            </transition>
          </div>
        </div>
      </div>
    </div>
    <!-- 黑名单部分 -->
    <div class="blacklist">
      <dv-border-box-4>
        <div class="blacklist-title">黑名单</div>
        <div class="blacklist-body">
          <el-row>
            <el-col :span="24">
              <el-card shadow="always">
                ID:1674289 用户名:心花小子 新浪微博 谣言发布次数:65
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:1243239 用户名:百事可爱 新浪微博 谣言发布次数:57
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:1567434 用户名:寄云间 新浪微博 谣言发布次数:49
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:193234534 用户名:鸭梨山大 b站 谣言发布次数:43
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:1634572 用户名:呦呦儿 新浪微博 谣言发布次数:33
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:182589234 用户名:隔壁小孩 b站 谣言发布次数:31
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:1272345 用户名:怀中猫 新浪微博 谣言发布次数:29
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:1678025798 用户名:空山梦 b站 谣言发布次数:27
              </el-card>
            </el-col>
          </el-row>
        </div>
      </dv-border-box-4>
    </div>
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
import * as CommentData from '../utils/mockdata'
import { getRumourInfo } from '@/utils/api'

export default {
  name: 'dataSquare',
  components: {
    Header
  },
  data() {
    return {
      textarea: '',
      comments: [],
      inputComment: '',
      showItemId: '',
      targetName: ''
    }
  },
  mounted() {
    this.comments = CommentData.comment.data
    // this.getInfo()
  },
  methods: {
    ChangeLike(item, type) {
      item.likeStatus = type
      if (type == 0) {
        // 取消点赞
        item.likeNum--
      } else if (type == 1) {
        //点赞
        item.likeNum++
      }
    },
    showCommentInput(item, reply) {
      if (reply) {
        this.inputComment = '@' + reply.fromName + ' '
        this.targetName = reply.fromName + ' '
      } else {
        this.inputComment = ''
        this.targetName = item.fromName
      }
      this.showItemId = item.id
    },
    cancelComment() {
      this.showItemId = ''
    },
    cancelSearch() {
      console.log('清空输入框')
      this.textarea = ''
      this.commitSearch()
    },
    getNowTime() {
      let year = new Date().getFullYear() //获取当前时间的年份
      let month = new Date().getMonth() + 1 //获取当前时间的月份
      let day = new Date().getDate() //获取当前时间的天数
      let hours = new Date().getHours() //获取当前时间的小时
      let minutes = new Date().getMinutes() //获取当前时间的分数
      //当小于 10 的是时候，在前面加 0
      if (hours < 10) {
        hours = '0' + hours
      }
      if (minutes < 10) {
        minutes = '0' + minutes
      }
      //拼接格式化当前时间
      let times = year + '-' + month + '-' + day + ' ' + hours + ':' + minutes
      return times
    },
    commitComment(item) {
      console.log('评论内容：', this.inputComment)
      console.log(this.comments)
      let content = ''
      if (this.inputComment.startsWith('@')) {
        content = this.inputComment.split(' ').slice(1).join(' ')
      } else {
        content = this.inputComment
      }
      console.log(content)
      let newComment = {
        id: item.id, //主键id
        commentId: item.commentId, //父评论id，即父亲的id
        fromId: 'observer2232432', //评论者id
        fromName: '枯城', //评论者昵称
        toId: item.commentId, //被评论者id
        toName: this.targetName, //被评论者昵称
        content: content, //评论内容
        date: this.getNowTime() //评论时间
      }
      item.reply.push(newComment)
      this.showItemId = ''
    },
    commitSearch() {
      console.log('动态内容:', this.textarea)
      this.comments.forEach((item) => {
        if (item.content.includes(this.textarea)) {
          item.showStatus = true
        } else {
          item.showStatus = false
        }
      })
    },
    changeAgree(item, type) {
      item.agree = type
    },
    getInfo() {
      let data = {
        id: this.$store.state.userInfo.data.id
      }
      getRumourInfo(data)
        .then((res) => {
          console.log('获取推送:', res)
          if (res.status == 200) {
            console.log('成功')
          } else {
            console.log('失败')
            this.$Message.error(res.msg)
          }
        })
        .catch(() => {
          this.isLoading = false
        })
    }
  }
}
</script>

<style lang="less">
.body-container {
  background: url('../assets/footer-animation-bg.svg') center center;
  background-color: #050828;
}

.header {
  height: 106px;
  justify-content: space-between;
  align-items: center;
}

.info-section {
  display: inline-block;
  width: 1150px;

  .info-header {
    margin-left: 25%;
    padding-top: 20px;
    margin-bottom: 10px;
    width: 768px;
    background-color: #363b79;
    border-radius: 10px;

    .btn-control {
      display: flex;
      justify-content: flex-end;
      align-items: center;
      padding-bottom: 10px;

      padding-top: 10px;

      .cancel {
        font-size: 16px;
        margin-right: 20px;
        cursor: pointer;
        color: white;
      }
      .btn {
        margin-right: 20px;
        font-size: 16px;
        background-color: rgb(48, 75, 210);
      }
    }

    .el-textarea {
      margin-left: 10px;
      margin-right: 10px;
      width: calc(100% - 20px);
    }

    .el-textarea__inner {
      font-size: 20px;
      color: white;
      background-color: #2e3377;
      border: none;
      border-radius: 20px;
    }
  }

  .info-body {
    margin-left: 25%;
    padding-top: 20px;
    width: 768px;

    .info-show {
      width: 100%;
      margin-bottom: 20px;
      font-size: 20px;
      color: white;
      background-color: #363b79;
      border: none;
      border-radius: 10px;
    }
    .info-show:hover {
      box-shadow: 0 7px 15px 0 rgba(49, 196, 190, 0.75);
    }
  }
}

.blacklist {
  display: inline-block;
  width: 600px;
  height: 700px;
  color: white;
  text-align: center;
  position: absolute;
  margin-top: 50px;

  .dv-border-box-4 {
    .blacklist-title {
      font-size: 25px;
      padding-top: 40px;
    }

    .blacklist-body {
      height: 585px;
      overflow-y: auto;

      .el-col-24 {
        margin-top: 20px;
        margin-left: 10px;
        padding-left: 10px;
        padding-right: 10px;
        width: calc(100% - 20px);
      }

      .el-card:hover {
        background-image: linear-gradient(to right, #6253e1, #852d91);
        box-shadow: 0 7px 15px 0 rgba(236, 116, 149, 0.75);
      }

      .el-card {
        font-size: 20px;
        color: white;
        background-color: #363b79;
        border: none;
        border-radius: 20px;
      }

      .el-card:hover {
        background-image: linear-gradient(to right, #6253e1, #852d91);
        box-shadow: 0 7px 15px 0 rgba(236, 116, 149, 0.75);
      }
    }

    .blacklist-body::-webkit-scrollbar {
      width: 10px;
      height: 10px;
      /**/
    }
    .blacklist-body::-webkit-scrollbar-track {
      background: rgb(239, 239, 239);
      border-radius: 2px;
    }
    .blacklist-body::-webkit-scrollbar-thumb {
      background: #bfbfbf;
      border-radius: 10px;
    }
    .blacklist-body::-webkit-scrollbar-thumb:hover {
      background: #333;
    }
    .blacklist-body::-webkit-scrollbar-corner {
      background: #179a16;
    }
  }
}

.info-body {
  .info-show {
    padding: 10px;

    .info {
      .info-name {
        font-size: 30px;
        color: skyblue;
        font-weight: 900;
      }

      .info-date {
        margin-left: 20px;
        font-size: 15px;
        color: rgb(142, 144, 158);
      }
    }

    .content {
      margin-top: 10px;
      font-size: 20px;
      color: white;
      background-color: #2e3377;
      border: none;
      border-radius: 20px;

      .content-about {
        text-indent: 2em;
        padding-top: 10px;
        padding-bottom: 10px;
        margin-left: 10px;
      }
    }

    .control {
      margin-top: 10px;
      height: 30px;

      .control-dianzan {
        height: 30px;
        margin-left: 20px;

        img {
          margin-top: -8px;
          width: 30px;
          height: 30px;
        }
        span {
          margin-left: 5px;
        }
      }
      .control-pinglun {
        margin-left: 20px;

        img {
          margin-top: -4px;
          width: 30px;
          height: 30px;
        }
      }

      .control-agree {
        margin-left: 40px;
      }

      .control-disagree {
        margin-left: 20px;
      }

      .control-agree:hover,
      .control-disagree:hover,
      .control-dianzan:hover,
      .control-pinglun:hover {
        cursor: pointer;
      }
    }

    .reply {
      .item {
        margin-top: 15px;
        font-size: 16px;
        color: white;
        background-color: #2e3377;
        border: none;
        border-radius: 20px;

        .reply-content {
          text-indent: 2em;
          padding-top: 10px;
          margin-left: 10px;

          .from-name,
          .to-name {
            color: skyblue;
          }
        }

        .reply-bottom {
          padding-bottom: 10px;
          .reply-bottom-date {
            margin-left: 50px;
            margin-bottom: 30px;
            font-size: 12px;
            color: rgb(142, 144, 158);
          }
          .reply-text {
            margin-left: 15px;
            img {
              width: 20px;
              height: 20px;
            }
          }
        }

        .reply-bottom:hover {
          cursor: pointer;
        }
      }

      .input-wrapper {
        margin-top: 20px;

        .el-textarea__inner {
          font-size: 16px;
          color: white;
          background-color: #2e3377;
          border: none;
          border-radius: 20px;
        }

        .btn-control {
          display: flex;
          justify-content: flex-end;
          align-items: center;

          padding-top: 10px;

          .cancel {
            font-size: 16px;
            margin-right: 20px;
            cursor: pointer;
          }
          .btn {
            margin-right: 20px;
            font-size: 16px;
            background-color: rgb(48, 75, 210);
          }
        }
      }
    }
  }
}
</style>