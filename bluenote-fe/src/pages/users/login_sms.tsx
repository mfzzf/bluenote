import React, { useState } from 'react';
import { Button, Form, Input, Card, Space } from 'antd';
import axios from "@/axios/axios";
import router from "next/router";
import Link from "next/link";
import '../../styles/global.css';

const onFinish = (values: any) => {
    axios.post("/users/login_sms", values)
        .then((res) => {
            if(res.status != 200) {
                alert(res.statusText);
                return
            }

            if (res.data.code == 0) {
                router.push('/users/profile')
                return;
            }
            alert(res.data.msg)
        }).catch((err) => {
        alert(err);
    })
};

const onFinishFailed = (errorInfo: any) => {
    alert("输入有误")
};

const LoginFormSMS: React.FC = () => {
    const [form] = Form.useForm();
    const [countdown, setCountdown] = useState(0);

    const sendCode = () => {
        const phone = form.getFieldValue("phone");
        
        if (!phone) {
            alert("请输入手机号码");
            return;
        }
        
        axios.post("/users/login_sms/code/send", {"phone": phone}).then((res) => {
            if(res.status != 200) {
                alert(res.statusText);
                return;
            }
            
            // Start countdown
            setCountdown(60);
            const timer = setInterval(() => {
                setCountdown(prevCount => {
                    if (prevCount <= 1) {
                        clearInterval(timer);
                        return 0;
                    }
                    return prevCount - 1;
                });
            }, 1000);
            
            alert(res?.data?.msg || "验证码已发送");
        }).catch((err) => {
            alert(err);
        });
    };

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
                    }}>手机号登录</h1>
                    <p style={{ color: '#666', marginTop: '8px' }}>使用手机号快速登录</p>
                </div>
                
                <Form
                    name="sms_login"
                    layout="vertical"
                    form={form}
                    onFinish={onFinish}
                    onFinishFailed={onFinishFailed}
                    autoComplete="off"
                >
                    <Form.Item
                        name="phone"
                        rules={[{ required: true, message: '请输入手机号码' }]}
                    >
                        <Input
                            className="redbook-input"
                            placeholder="请输入手机号码"
                            size="large"
                        />
                    </Form.Item>

                    <Form.Item
                        name="code"
                        rules={[{ required: true, message: '请输入验证码' }]}
                    >
                        <div style={{ display: 'flex', gap: '10px' }}>
                            <Input
                                className="redbook-input"
                                placeholder="请输入验证码"
                                size="large"
                                style={{ flex: 1 }}
                            />
                            <Button
                                style={{
                                    height: '44px',
                                    borderRadius: '8px',
                                    width: '120px',
                                    fontWeight: '500'
                                }}
                                disabled={countdown > 0}
                                onClick={sendCode}
                            >
                                {countdown > 0 ? `${countdown}秒后重试` : '发送验证码'}
                            </Button>
                        </div>
                    </Form.Item>

                    <Form.Item style={{ marginBottom: '12px' }}>
                        <Button className="redbook-button" type="primary" htmlType="submit" block>
                            登录/注册
                        </Button>
                    </Form.Item>
                    
                    <div style={{ 
                        textAlign: 'center',
                        marginTop: '20px',
                        borderTop: '1px solid #f0f0f0',
                        paddingTop: '20px'
                    }}>
                        <Space size={30} style={{ color: '#666' }}>
                            <Link href="/users/login" style={{ color: '#666' }}>邮箱登录</Link>
                            <Link href="/users/login_wechat" style={{ color: '#666' }}>微信登录</Link>
                            <Link href="/users/signup" style={{ color: 'var(--primary-color)' }}>注册账号</Link>
                        </Space>
                    </div>
                </Form>
            </Card>
        </div>
    );
};

export default LoginFormSMS;