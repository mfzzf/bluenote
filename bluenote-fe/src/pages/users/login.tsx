import React, { useState } from 'react';
import { Button, Form, Input, Card, Space, message, Spin } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import axios from "@/axios/axios";
import Link from "next/link";
import router from "next/router";
import '../../styles/variables.css';
import styles from '../../styles/login.module.css';

const LoginForm: React.FC = () => {
    const [loading, setLoading] = useState(false);
    
    const onFinish = async (values: any) => {
        setLoading(true);
        try {
            const res = await axios.post("/users/login", values);
            if (res.status === 200) {
                message.success('登录成功');
                router.push('/users/profile');
            } else {
                message.error(res.statusText || '登录失败');
            }
        } catch (err: any) {
            message.error(err.message || '登录失败，请稍后再试');
        } finally {
            setLoading(false);
        }
    };

    const onFinishFailed = () => {
        message.error("请检查输入信息");
    };

    return (
        <div className={styles.loginContainer}>
            <Card className={styles.loginCard}>
                <div className={styles.headerContainer}>
                    <h1 className={styles.gradientTitle}>欢迎回来</h1>
                    <p className={styles.subtitle}>登录你的账号</p>
                </div>
                
                <Form
                    name="login"
                    layout="vertical"
                    onFinish={onFinish}
                    onFinishFailed={onFinishFailed}
                    autoComplete="off"
                >
                    <Form.Item
                        name="email"
                        rules={[
                            { required: true, message: '请输入邮箱' },
                            { type: 'email', message: '请输入有效的邮箱地址' }
                        ]}
                    >
                        <Input 
                            prefix={<UserOutlined className={styles.inputIcon} />}
                            className={styles.formInput}
                            placeholder="请输入邮箱"
                            size="large" 
                        />
                    </Form.Item>

                    <Form.Item
                        name="password"
                        rules={[{ required: true, message: '请输入密码' }]}
                    >
                        <Input.Password 
                            prefix={<LockOutlined className={styles.inputIcon} />}
                            className={styles.formInput}
                            placeholder="请输入密码"
                            size="large" 
                        />
                    </Form.Item>

                    <Form.Item style={{ marginBottom: '12px' }}>
                        <Button 
                            className={styles.loginButton} 
                            type="primary" 
                            htmlType="submit" 
                            block
                            loading={loading}
                        >
                            登录
                        </Button>
                    </Form.Item>

                    <div className={styles.footerLinks}>
                        <Space size={30}>
                            <Link href="/users/login_sms" className={styles.normalLink}>手机号登录</Link>
                            <Link href="/users/login_wechat" className={styles.normalLink}>微信登录</Link>
                            <Link href="/users/signup" className={styles.highlightLink}>注册账号</Link>
                        </Space>
                    </div>
                </Form>
            </Card>
        </div>
    );
};

export default LoginForm;