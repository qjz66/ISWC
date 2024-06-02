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
          :rows="4"
          placeholder="快来发表动态吧！"
          v-model="textarea"
        >
        </el-input>
        <el-button type="primary" @click="commitUpdate">发布</el-button>
      </div>
      <!-- 动态显示 -->
      <div class="info-body" v-for="item in comments" :key="item.date">
        <div class="info-show">
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
                  <span class="cancel" @click="cancel">取消</span>
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
                ID:001 用户名:张三 新浪微博 谣言发布次数:12
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:001 用户名:张三 新浪微博 谣言发布次数:12
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:001 用户名:张三 新浪微博 谣言发布次数:12
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:001 用户名:张三 新浪微博 谣言发布次数:12
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:001 用户名:张三 新浪微博 谣言发布次数:12
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:001 用户名:张三 新浪微博 谣言发布次数:12
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:001 用户名:张三 新浪微博 谣言发布次数:12
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:001 用户名:张三 新浪微博 谣言发布次数:12
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:001 用户名:张三 新浪微博 谣言发布次数:12
              </el-card>
            </el-col>
            <el-col :span="24">
              <el-card shadow="always">
                ID:001 用户名:张三 新浪微博 谣言发布次数:12
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
      showItemId: ''
    }
  },
  mounted() {
    this.comments = CommentData.comment.data
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
      } else {
        this.inputComment = ''
      }
      this.showItemId = item.id
    },
    cancel() {
      this.showItemId = ''
    },
    commitComment(item) {
      console.log('评论内容：', this.inputComment)
      console.log(this.comments)
      let newComment = {
        id: '345232144545', //主键id
        commentId: 'comment20001', //父评论id，即父亲的id
        fromId: 'observer2232432', //评论者id
        fromName: '夕阳红', //评论者昵称
        toId: 'errhefe2323213', //被评论者id
        toName: '犀利的评论家', //被评论者昵称
        content: this.inputComment, //评论内容
        date: '2018-37-05 08:35' //评论时间
      }
      item.reply.push(newComment)
    },
    commitUpdate() {
      console.log('动态内容:', this.textarea)
      let newUpdate = {
        id: 'comment01002',
        date: '2018-17-15 08:30',
        ownerId: 'talents2100020',
        fromId: 'errhefe2322213',
        fromName: '毒蛇郭德纲111',
        likeNum: 0,
        likeStatus: 0,
        content: this.textarea,
        reply: []
      }
      this.comments.push(newUpdate)
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

    .el-button {
      margin-top: 15px;
      margin-left: 20px;
      margin-bottom: 15px;
      border-radius: 15px;
      border: none;
      font-size: 16px;
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

      .el-card {
        font-size: 20px;
        color: white;
        background-color: #363b79;
        border: none;
        border-radius: 20px;
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
          }
        }
      }
    }
  }
}
</style>