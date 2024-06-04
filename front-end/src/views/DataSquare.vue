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
      <div class="info-body" v-for="item in comments" :key="item.id">
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
                @click="changeLike(item, 1)"
                alt=""
              />
              <span>{{ item.likeNum }}</span>
            </span>
            <span class="control-dianzan" v-else-if="item.likeStatus === 1">
              <img
                src="../picture/dianzan_red.png"
                @click="changeLike(item, 0)"
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
            <div class="item" v-for="reply in item.comment" :key="reply.id">
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
            <el-col :span="24" v-for="black in blackList" :key="black.id">
              <el-card shadow="always">
                ID:{{ black.id }} 用户名:{{ black.username }}
                {{ black.platform }} 谣言发布次数:{{ black.number }}
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

import {
  getRumourInfo,
  getNowTime,
  getBlackList,
  showLike,
  postCommit
} from '@/utils/api'

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
      targetName: '',
      targetId: '',
      blackList: []
    }
  },
  mounted() {
    // this.comments = CommentData.comment.data
    this.getCommentInfo()
    this.getBlackListInfo()
  },
  methods: {
    // 搜索动态
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
    // 取消搜索,即清空搜索内容
    cancelSearch() {
      console.log('清空输入框')
      this.textarea = ''
      this.commitSearch()
    },
    // 点赞或者取消点赞
    changeLike(item, type) {
      item.likeStatus = type
      if (type == 0) {
        // 取消点赞
        item.likeNum--
      } else if (type == 1) {
        //点赞
        item.likeNum++
      }

      let params = {
        id: this.$store.state.userInfo.data.id
      }

      let data = new FormData()
      data.append('updateId', item.id)

      showLike(params, data).then((res) => {
        console.log('点赞:', res)
        if (res.status == 200) {
          console.log('点赞成功')
        } else {
          console.log('点赞失败')
          this.$message.error(res.msg)
        }
      })
    },
    // 展示出评论输入框
    showCommentInput(item, reply) {
      if (reply) {
        this.inputComment = '@' + reply.fromName + ' '
        this.targetName = reply.fromName + ' '
        this.targetId = reply.fromId
      } else {
        this.inputComment = ''
        this.targetName = item.fromName
        this.targetId = item.fromId
      }
      this.showItemId = item.id
    },
    // 取消评论(隐去评论输入框)
    cancelComment() {
      this.showItemId = ''
    },
    // 提交评论
    commitComment(item) {
      // 格式化输入的内容
      console.log('评论内容：', this.inputComment)
      console.log(this.comments)
      let content = ''
      if (this.inputComment.startsWith('@')) {
        content = this.inputComment.split(' ').slice(1).join(' ')
      } else {
        content = this.inputComment
      }
      console.log(content)

      let params = {
        id: this.$store.state.userInfo.data.id
      }

      let data = new FormData()
      data.append('updateId', item.id)
      data.append('fromName', this.$store.state.userInfo.data.username)
      data.append('toId', this.targetId)
      data.append('toName', this.targetName)
      data.append('content', content)
      data.append('date', getNowTime())

      postCommit(params, data).then((res) => {
        console.log('评论:', res)
        if (res.status == 200) {
          console.log('评论成功')

          let newComment = {
            id: res.data.commentId, //主键id
            updateId: item.id, //父评论id，即父亲的id
            fromId: this.$store.state.userInfo.data.id, //评论者id
            fromName: this.$store.state.userInfo.data.userName, //评论者昵称
            toId: this.targetId, //被评论者id
            toName: this.targetName, //被评论者昵称
            content: content, //评论内容
            date: getNowTime() //评论时间
          }
          console.log(newComment)
          item.comment.push(newComment)

          this.showItemId = ''
        } else {
          console.log('评论失败')
          this.$message.error(res.msg)
        }
      })
    },
    // 赞成和反对
    changeAgree(item, type) {
      item.agree = type
    },
    // 获取动态列表以及评论列表、黑名单
    getCommentInfo() {
      let params = {
        id: this.$store.state.userInfo.data.id
      }
      // 获取动态列表以及评论列表
      getRumourInfo(params).then((res) => {
        console.log('获取推送:', res)
        if (res.status == 200) {
          console.log('获取推送成功')
          this.comments = res.data.updates
          console.log('this.comments', this.comments)
          console.log(typeof this.comments[0].comment)
        } else {
          console.log('获取推送失败')
          this.$message.error(res.msg)
        }
      })
    },
    getBlackListInfo() {
      let params = {
        id: this.$store.state.userInfo.data.id
      }
      // 获取黑名单
      getBlackList(params).then((res) => {
        console.log('获取黑名单:', res)
        if (res.status == 200) {
          console.log('获取黑名单成功')
          this.blackList = res.data.blacklist
        } else {
          console.log('获取黑名单失败')
          this.$message.error(res.msg)
        }
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