
const WebSocket = require('ws');

function testCollaboration() {
    const users = [
        { id: 1, username: '用户A' },
        { id: 2, username: '用户B' },
        { id: 3, username: '用户C' }
    ];
    
    const connections = [];
    const noteId = 1;
    
    console.log('开始测试协同编辑功能...');
    
    users.forEach((user, index) => {
        const ws = new WebSocket(`ws://localhost:8888/api/v1/collaboration/ws?noteId=${noteId}&userId=${user.id}&username=${user.username}`);
        
        ws.on('open', () => {
            console.log(`${user.username} 已连接`);
            connections.push({ ws, user });
            
            setTimeout(() => {
                sendOperation(ws, 'insert', index * 10, `来自${user.username}的内容`);
            }, index * 1000);
        });
        
        ws.on('message', (data) => {
            const message = JSON.parse(data);
            console.log(`${user.username} 收到消息:`, message.type);
            
            if (message.type === 'operation') {
                console.log(`${user.username} 收到操作:`, message.operation);
            }
        });
        
        ws.on('close', () => {
            console.log(`${user.username} 连接已关闭`);
        });
        
        ws.on('error', (error) => {
            console.error(`${user.username} 连接错误:`, error.message);
        });
    });
    
    setTimeout(() => {
        connections.forEach(({ ws }) => {
            ws.close();
        });
        console.log('测试完成');
    }, 10000);
}

function sendOperation(ws, type, position, content) {
    const message = {
        type: 'operation',
        opType: type,
        position: position,
        content: content
    };
    ws.send(JSON.stringify(message));
    console.log('发送操作:', message);
}

if (require.main === module) {
    testCollaboration();
}

module.exports = { testCollaboration }; 