//模拟评论数据
const comment = {
  status: "成功",
  code: 200,
  data: [
    {
      id: 'comment0001', //主键id
      date: '2024-6-1 15:00',  //评论时间
      fromId: 'errhefe232213',  //评论者id
      fromName: '评论家',   //评论者昵称
      likeNum: 3, //点赞人数
      likeStatus: 0,
      showStatus: true,
      agree: 2,
      content: '【#怀孕学生被当肾病医治存在漏诊误诊# 】#怀孕学生被当肾病医治已达成赔偿协议# 5月31日，“怀胎7月女大学生被当肾病医治后身亡”一事，引发社会关注。',  //评论内容
      reply: [  //回复，或子评论
        {
          id: '34523244545',  //主键id
          updateId: 'comment0001',  //父评论id，即父亲的id
          fromId: 'observer223432',  //评论者id
          fromName: '夕阳红',  //评论者昵称
          toId: 'errhefe232213',  //被评论者id
          toName: '评论家',  //被评论者昵称
          content: '新闻分析的准确性很高！这个医生怎么能把怀孕误检了',  //评论内容
          date: '2024-6-2 11:45'   //评论时间
        },
        {
          id: '34523244545',
          updateId: 'comment0001',
          fromId: 'observer567422',
          fromName: '清晨一缕阳光',
          toId: 'observer223432',
          toName: '夕阳红',
          content: '是的，这个医生太不负责了。',
          date: '2024-6-2 16:09'
        }
      ]
    },
    {
      id: 'comment0002',
      date: '2024-5-25 13:52',
      fromId: 'errhefe232213',
      fromName: '毒蛇纲',
      likeNum: 4,
      likeStatus: 0,
      showStatus: true,
      agree:2,
      content: '【#生长激素不能随意用#！#这些因素影响儿童身高#】据央视新闻报道：现在家长对孩子的身高越来越关注，一些家长为了让孩子长得更高，给孩子使用了生长激素。针对这一现象，国家儿童医学中心主任、北京儿童医院院长倪鑫表示，孩子的身高取决于两个因素，这两个因素70%是遗传因素，30%是后天因素，对于生长激素的运用，一定要经过医生专业诊治。',
      reply: [
        {
        id: '34523244545',  //主键id
        updateId: 'comment0001',  //父评论id，即父亲的id
        fromId: 'observer223432',  //评论者id
        fromName: '春天',  //评论者昵称
        toId: 'errhefe232213',  //被评论者id
        toName: '毒蛇纲',  //被评论者昵称
        content: '新闻分析的很不错，很值得借鉴！',  //评论内容
        date: '2024-5-27 11:45'   //评论时间
        }
      ]
    }
  ]
};

export {comment}
