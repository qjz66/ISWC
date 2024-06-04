import network from './network';

// 获取当前时间
export function getNowTime() {
  let year = new Date().getFullYear(); //获取当前时间的年份
  let month = new Date().getMonth() + 1; //获取当前时间的月份
  let day = new Date().getDate(); //获取当前时间的天数
  let hours = new Date().getHours(); //获取当前时间的小时
  let minutes = new Date().getMinutes(); //获取当前时间的分数
    //当小于 10 的是时候，在前面加 0
    if (hours < 10) {
      hours = '0' + hours
    }
    if (minutes < 10) {
      minutes = '0' + minutes
    }
    //拼接格式化当前时间
  let times = year + '-' + month + '-' + day + ' ' + hours + ':' + minutes;
  return times;
}

// 登录
export function login(data) {
  return network({
    url: `/rd/user/login`,
    method: "post",
    data
  });
}

// 注册
export function register(data) {
  return network({
    url: `/rd/user/register`,
    method: "post",
    data
  })
}

// 密码重置
export function resetPwd(data) {
  return network({
    url: `/resetPwd`,
    method: "post",
    data
  })
}

// 数据广场
// 获取动态列表以及评论列表
export function getRumourInfo(params) {
  return network({
    url: `/rd/ground/get_updates`,
    method: "get",
    params
  });
}

// 点赞
export function showLike(params, data) {
  return network({
    url: `/rd/ground/favorite`,
    method: "post",
    params,
    data
  });
}
// 评论
export function postCommit(params, data) {
  return network({
    url: `/rd/ground/comment`,
    method: "post",
    params,
    data
  });
}

// 获取黑名单
export function getBlackList(params) {
  return network({
    url: `/rd/ground/blacklist`,
    method: "get",
    params
  });
}

// 个人中心
// 获取历史记录
export function getHistory(params) {
  return network({
    url: `/rd/user/history`,
    method: "get",
    params
  });
}

// 检测谣言
// 进行检测/文本检测
export function analysisNews(params, data) {
  return network({
    url: `/rd/detect/upload_text`,
    method: "post",
    params,
    data
  });
}

// 分析结果
// 分享报告/上传动态
export function shareAnalysisResult(params, data) {
  return network({
    url: `/rd/ground/update`,
    method: "post",
    params,
    data
  });
}