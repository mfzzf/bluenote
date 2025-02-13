import React, {useEffect, useState} from 'react';
import { Button, Card, Spin } from 'antd';
import axios from "@/axios/axios";
import router from "next/router";
import { UserOutlined, MailOutlined, PhoneOutlined, CalendarOutlined } from '@ant-design/icons';
import '../../styles/profile.css';

function Profile() {
    const [data, setData] = useState<Profile | null>(null)
    const [isLoading, setLoading] = useState(false)

    useEffect(() => {
        setLoading(true)
        axios.get('/users/profile')
            .then((res) => res.data)
            .then((data) => {
                setData(data)
                setLoading(false)
            })
    }, [])

    if (isLoading) return (
        <div style={{ 
            display: 'flex', 
            justifyContent: 'center', 
            alignItems: 'center', 
            height: '100vh',
            background: 'linear-gradient(135deg, #fff1f1 0%, #fff 100%)'
        }}>
            <Spin size="large" />
        </div>
    )

    if (!data) return <p>No profile data</p>

    return (
        <div style={{ background: '#f5f5f5', minHeight: '100vh' }}>
            <div className="profile-header">
                <div className="profile-avatar">
                    <UserOutlined style={{ fontSize: '48px', color: '#ff2442', marginTop: '30px' }} />
                </div>
                <h1 style={{ 
                    fontSize: '24px',
                    margin: '0 0 8px 0',
                    color: '#333'
                }}>{data.Nickname || '未设置昵称'}</h1>
                <p style={{ color: '#666', margin: 0 }}>
                    {data.AboutMe || '这个人很懒，还没有写简介...'}
                </p>
            </div>

            <div style={{ padding: '0 20px', maxWidth: '800px', margin: '0 auto' }}>
                <div className="profile-info-card">
                    <div className="profile-info-item">
                        <span className="profile-info-label">
                            <MailOutlined /> 邮箱
                        </span>
                        <span className="profile-info-value">{data.Email}</span>
                    </div>
                    
                    <div className="profile-info-item">
                        <span className="profile-info-label">
                            <PhoneOutlined /> 手机号
                        </span>
                        <span className="profile-info-value">{data.Phone || '未绑定'}</span>
                    </div>
                    
                    <div className="profile-info-item">
                        <span className="profile-info-label">
                            <CalendarOutlined /> 生日
                        </span>
                        <span className="profile-info-value">{data.Birthday || '未设置'}</span>
                    </div>

                    <Button 
                        type="primary"
                        className="edit-button"
                        block
                        onClick={() => router.push('/users/edit')}
                    >
                        编辑资料
                    </Button>
                </div>
            </div>
        </div>
    )
}

export default Profile;