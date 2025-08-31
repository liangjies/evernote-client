export function friendlyDate(dateStr) {
  let dateObj = typeof dateStr === 'object' ? dateStr : new Date(dateStr)
  let time = dateObj.getTime()
  let now = Date.now()
  let space = now - time
  let str = ''

  switch (true) {
    case space < 1000 * 60:
      str = '刚刚'
      break
    case space < 1000 * 3600:
      str = Math.floor(space / (1000 * 60)) + '分钟前'
      break
    case space < 1000 * 3600 * 24:
      str = Math.floor(space / (1000 * 3600)) + '小时前'
      break
    default:
      let date = getFullDate(dateStr).split(' ')[0]
      
      var year = date.split('-')[0]+'年';
      var month = date.split('-')[1]+'月';
      var date = date.split('-')[2]+'日';
      str = year + month + date
      // str = Math.floor(space / (1000 * 3600 * 24)) + '天前'
      break
  }
  return str
}

export function getFullDate(dateStr) {
  // let lastDotIndex = dateStr.lastIndexOf('.')
  // dateStr = dateStr.substring(0, lastDotIndex)
  // return dateStr.split('T').join(' ')
  let newDate
  if(dateStr.indexOf("Z") == -1){
    var date = new Date(dateStr).toJSON();
    newDate=new Date(+new Date(date)+8*3600*1000).toISOString().replace(/T/g,' ').replace(/\.[\d]{3}Z/,'');
  }else{
    let lastDotIndex = dateStr.lastIndexOf('.')
    dateStr = dateStr.substring(0, lastDotIndex)
    newDate = dateStr.split('T').join(' ')
  }

  return newDate
}
