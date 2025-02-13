import React from 'react';
import { Button, Form, Input, Card, Space } from 'antd';
import axios from "@/axios/axios";
import Link from "next/link";
import router from "next/router";

const onFinish = (values: any) => {
    axios.post("/users/login", values)
        .then((res) => {
            if(res.status != 200) {
                alert(res.statusText);
                return
            }
            alert(res.data)
            router.push('/users/profile')
        }).catch((err) => {
            alert(err);
    })
};

const onFinishFailed = (errorInfo: any) => {
    alert("输入有误")
};

const LoginForm: React.FC = () => {
    return (
        <div style={{ 
            display: 'flex', 
            justifyContent: 'center', 
            alignItems: 'center', 
            minHeight: '100vh',
            background: 'linear-gradient(135deg, #fff1f1 0%, #fff 100%)',
            padding: '20px'
        }}>
            <Card 
                className="redbook-card"
                style={{ 
                    width: '100%',
                    maxWidth: 400,
                    border: 'none'
                }}
            >
                <div style={{
                    textAlign: 'center',
                    marginBottom: '30px'
                }}>
                    <h1 style={{
                        fontSize: '28px',
                        background: 'var(--background-gradient)',
                        WebkitBackgroundClip: 'text',
                        WebkitTextFillColor: 'transparent',
                        margin: 0
                    }}>欢迎回来</h1>
                    <p style={{ color: '#666', marginTop: '8px' }}>登录你的账号</p>
                </div>
                
                <Form
                    name="basic"
                    layout="vertical"
                    onFinish={onFinish}
                    onFinishFailed={onFinishFailed}
                    autoComplete="off"
                >
                    <Form.Item
                        name="email"
                        rules={[{ required: true, message: '请输入邮箱' }]}
                    >
                        <Input 
                            className="redbook-input"
                            placeholder="请输入邮箱"
                            size="large" 
                        />
                    </Form.Item>

                    <Form.Item
                        name="password"
                        rules={[{ required: true, message: '请输入密码' }]}
                    >
                        <Input.Password 
                            className="redbook-input"
                            placeholder="请输入密码"
                            size="large" 
                        />
                    </Form.Item>

                    <Form.Item style={{ marginBottom: '12px' }}>
                        <Button className="redbook-button" type="primary" htmlType="submit" block>
                            登录
                        </Button>
                    </Form.Item>

                    <div style={{ 
                        textAlign: 'center',
                        marginTop: '20px',
                        borderTop: '1px solid #f0f0f0',
                        paddingTop: '20px'
                    }}>
                        <Space size={30} style={{ color: '#666' }}>
                            <Link href="/users/login_sms" style={{ color: '#666' }}>手机号登录</Link>
                            <Link href="/users/login_wechat" style={{ color: '#666' }}>微信登录</Link>
                            <Link href="/users/signup" style={{ color: 'var(--primary-color)' }}>注册账号</Link>
                        </Space>
                    </div>
                </Form>
            </Card>
        </div>
    );
};

export default LoginForm;