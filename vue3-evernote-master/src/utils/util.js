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

            let year = date.split('-')[0] + '/';
            let month = date.split('-')[1] + '/';
            let day = date.split('-')[2];
            str = year + month + day
            // str = Math.floor(space / (1000 * 3600 * 24)) + '天前'
            break
    }
    return str
}

export function getFullDate(dateStr) {
    // let lastDotIndex = dateStr.lastIndexOf('.')
    // dateStr = dateStr.substring(0, lastDotIndex)
    // return dateStr.split('T').join(' ')
    if (typeof dateStr == "undefined") {
        return ""
    }
    let newDate
    if (dateStr.indexOf("Z") == -1) {
        let date = new Date(dateStr).toJSON();
        newDate = new Date(+new Date(date) + 8 * 3600 * 1000).toISOString().replace(/T/g, ' ').replace(/\.[\d]{3}Z/, '');
    } else {
        let lastDotIndex = dateStr.lastIndexOf('.')
        dateStr = dateStr.substring(0, lastDotIndex)
        newDate = dateStr.split('T').join(' ')
    }

    return newDate
}


export function isHtmlLabel(text, keyWord) {
    let html = '<a><abbr><acronym><address><applet><area><article><aside><audio><b><base><basefont><bdi><bdo><big><blockquote><body><br><button><canvas><caption><center><cite><code><col><colgroup><command><datalist><dd><del><details><dfn><dialog><dir><div><dl><dt><em><embed><fieldset><figcaption><figure><font><footer><form><frame><frameset><h1-h6><head><header><hr><html><i><iframe><img><input><ins><kbd><keygen><label><legend><li><link><map><mark><menu><meta><meter><nav><noframes><noscript><object><ol><optgroup><option><output><p><param><pre><progress><q><rp><rt><ruby><s><samp><script><section><select><small><source><span><strike><strong><style><sub><summary><sup><table><tbody><td><textarea><tfoot><th><thead><time><title><tr><track><tt><u><ul><var><video><wbr>'
    var re = new RegExp("<[a-zA-Z\/]+.*?>", "gi");
    let textA = text.match(re)
    if (textA == null) {
        return false
    }
    if (html.includes(keyWord)) {
        for (var i = 0; i < textA.length; i++) {
            if (textA[i].includes(keyWord)) {
                return true
            }
        }
    }
    return false
}