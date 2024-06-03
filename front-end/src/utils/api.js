import network from './network';

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

// 获取推送
export function getRumourInfo() {
  return network({
    url: `/rd/push/get_rumors`,
    method: "get",
  });
}
