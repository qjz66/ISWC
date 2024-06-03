<template>
  <div style="height: 100%" class="body-container">
    <div class="header">
      <Header />
    </div>
    <div class="personal-body">
      <!-- 基本信息 -->
      <div class="personal-info">
        <!-- 信息展示 -->
        <el-card>
          <el-descriptions
            class="margin-top"
            title="基本信息"
            :column="1"
            border
          >
            <template slot="extra">
              <el-button
                type="primary"
                size="small"
                @click="dialogPasswordVisible = true"
                >修改密码</el-button
              >
              <el-button
                type="primary"
                size="small"
                @click="dialogFormVisible = true"
                >修改信息</el-button
              >
            </template>
            <el-descriptions-item>
              <template slot="label">
                <i class="el-icon-user"></i>
                用户名
              </template>
              {{ info.userName }}
            </el-descriptions-item>
            <el-descriptions-item>
              <template slot="label">
                <i class="el-icon-odometer"></i>
                年龄
              </template>
              {{ info.age }}
            </el-descriptions-item>
            <el-descriptions-item>
              <template slot="label">
                <i class="el-icon-male"></i>
                <i class="el-icon-female"></i>
                性别
              </template>
              {{ info.sex }}
            </el-descriptions-item>
            <el-descriptions-item>
              <template slot="label">
                <i class="el-icon-message"></i>
                邮箱Email
              </template>
              {{ info.email }}
            </el-descriptions-item>
            <el-descriptions-item>
              <template slot="label">
                <i class="el-icon-mobile-phone"></i>
                手机号码
              </template>
              {{ info.phoneNumber }}
            </el-descriptions-item>
          </el-descriptions>
        </el-card>
        <!-- 密码修改 -->
        <el-dialog title="修改密码" :visible.sync="dialogPasswordVisible">
          <el-form :model="pwd">
            <el-form-item label="原密码" :label-width="formLabelWidth">
              <el-input
                type="password"
                v-model="pwd.oldPwd"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item label="新密码" :label-width="formLabelWidth">
              <el-input
                type="password"
                v-model="pwd.newPwd"
                autocomplete="off"
              ></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取 消</el-button>
            <el-button type="primary" @click="dialogFormVisible = false"
              >确 定</el-button
            >
          </div>
        </el-dialog>
        <!-- 信息修改 -->
        <el-dialog title="修改信息" :visible.sync="dialogFormVisible">
          <el-form :model="info">
            <el-form-item label="用户名" :label-width="formLabelWidth">
              <el-input v-model="info.userName" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="年龄" :label-width="formLabelWidth">
              <el-input v-model="info.age" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="性别" :label-width="formLabelWidth">
              <el-select v-model="info.sex" placeholder="请选择性别">
                <el-option label="男" value="男"></el-option>
                <el-option label="女" value="女"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="邮箱Email" :label-width="formLabelWidth">
              <el-input v-model="info.email" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="手机号码" :label-width="formLabelWidth">
              <el-input
                v-model="info.phoneNumber"
                autocomplete="off"
              ></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取 消</el-button>
            <el-button type="primary" @click="dialogFormVisible = false"
              >确 定</el-button
            >
          </div>
        </el-dialog>
      </div>
      <!-- 鉴别谣言的历史记录 -->
      <div class="personal-analysis">
        <dv-border-box-9 style="height: 100%; width: 100%">
          <div class="personal-title">鉴别谣言的历史记录</div>
          <div class="personal-list">
            <el-col :span="8">
              <el-card shadow="always">
                #20岁女生留学新加坡坠亡#，家属：离毕业不到10天，财物遗失。
              </el-card>
            </el-col>
            <el-col :span="8">
              <el-card shadow="always">
                #知名配音演员公开维权#被侵权配音演员林景称“真正让人在意的不是AI，而是侵权。”#云中君钟离配音演员遭AI侵权
              </el-card>
            </el-col>
            <el-col :span="8">
              <el-card shadow="always">
                【#怀孕学生被当肾病医治存在漏诊误诊#
                】#怀孕学生被当肾病医治已达成赔偿协议#
                5月31日，“怀胎7月女大学生被当肾病医治后身亡”一事，引发社会关注。
              </el-card>
            </el-col>
            <el-col :span="8">
              <el-card shadow="always">
                #白开水在4小时内饮用活性最佳#水对人体很重要，人可以三天无饭，但绝不能三天无水。想要健康，必须要正确认识水，选对水，喝对水。#矿泉水苏打水喝了真的健康吗#？
              </el-card>
            </el-col>
            <el-col :span="8">
              <el-card shadow="always">
                #顾客称一点点奶茶喝出壁虎#
                当事人称，他在涉事门店点了一杯布丁奶茶，快喝完时发现的壁虎。其称自己当时全程都在店里，没有离开过，店家查了监控也知道这不可能是他放进去的。对此，#一点点门店回应奶茶喝出壁虎#表示，是不可能存在这种情况的，有人正在联系对方处理此事。#顾客在奶茶里喝出整只壁虎#
              </el-card>
            </el-col>
          </div>
        </dv-border-box-9>
      </div>
    </div>
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
export default {
  name: 'dataSquare',
  components: {
    Header
  },
  data() {
    return {
      info: {
        userName: String,
        age: Number,
        email: String,
        phoneNumber: String,
        sex: String
      },
      pwd: {
        oldPwd: '',
        newPwd: ''
      },
      dialogPasswordVisible: false,
      dialogFormVisible: false,
      formLabelWidth: '120px'
    }
  },
  mounted() {
    this.load()
  },
  methods: {
    load() {
      console.log(111)
      this.info.userName = '枯城'
      this.info.age = 21
      this.info.email = '2831709830@qq.com'
      this.info.phoneNumber = '15570670729'
      this.info.sex = '男'
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

.personal-body {
  font-size: 20px;

  .personal-title {
    display: flex;
    justify-content: center;
    align-items: center;

    .dv-decoration-7 {
      background-image: linear-gradient(
        267deg,
        #641ff9 27.93%,
        #02c5f7 99.79%,
        #c6bafd -31.69%,
        #b7a2fd -22.78%
      );
      background-clip: text;
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      font-size: 45px;
    }
  }

  .personal-info {
    margin-top: 100px;
    margin-left: 400px;
    width: 400px;
    text-align: center;
    color: white;

    .el-descriptions__title {
      color: white;
    }

    .el-card {
      background-color: #1b226b;
      border: none;
      color: white;

      .el-descriptions-item__label.is-bordered-label {
        background-color: #1b226b;
        color: white;
      }

      .el-descriptions__body {
        background-color: #1b226b;
        color: white;
      }
    }

    img {
      width: 200px;
      height: 200px;
      object-fit: cover;
    }
  }

  .personal-analysis {
    float: right;
    height: 600px;
    margin-top: -330px;
    margin-right: 400px;
    width: 600px;

    text-align: center;
    color: white;

    .personal-title {
      padding-top: 25px;
      font-size: 25px;
    }

    .personal-list {
      width: 586px;
      height: 510px;
      overflow-y: auto;

      .el-col-8 {
        cursor: pointer;
        margin-left: 20px;
        padding-left: 20px;
        padding-right: 20px;
        width: calc(100% - 40px);
      }
      .el-card {
        font-size: 20px;
        color: white;
        background-color: #363b79;
        border: none;
        border-radius: 20px;
        margin-top: 20px;
      }

      .el-card:hover {
        background-image: linear-gradient(to right, #25aae1, #40e495);
        box-shadow: 0 7px 15px 0 rgba(49, 196, 190, 0.75);
      }
    }

    .personal-list::-webkit-scrollbar {
      width: 10px;
      height: 10px;
      /**/
    }
    .personal-list::-webkit-scrollbar-track {
      background: rgb(239, 239, 239);
      border-radius: 2px;
    }
    .personal-list::-webkit-scrollbar-thumb {
      background: #bfbfbf;
      border-radius: 10px;
    }
    .personal-list::-webkit-scrollbar-thumb:hover {
      background: #333;
    }
    .personal-list::-webkit-scrollbar-corner {
      background: #179a16;
    }
  }
}
</style>