import React, {useEffect, useState} from 'react';
import {Button, DatePicker, Form, Input, Card, Spin} from 'antd';
import axios from "@/axios/axios";
import moment from 'moment';
import router from "next/router";

const { TextArea } = Input;

const onFinish = (values: any) => {
    if (values.birthday) {
        values.birthday = moment(values.birthday).format("YYYY-MM-DD")
    }
    axios.post("/users/edit", values)
        .then((res) => {
            if(res.status != 200) {
                alert(res.statusText);
                return
            }
            if (res.data?.code == 0) {
                router.push('/users/profile')
                return
            }
            alert(res.data?.msg || "系统错误");
        }).catch((err) => {
        alert(err);
    })
};

const onFinishFailed = (errorInfo: any) => {
    alert("输入有误")
};

function EditForm() {
    let p: Profile = {Email: "", Phone: "", Nickname: "", Birthday:"", AboutMe: ""}
    const [data, setData] = useState<Profile>(p)
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
    );

    return (
        <div style={{ 
            padding: '24px',
            background: 'linear-gradient(135deg, #fff1f1 0%, #fff 100%)',
            minHeight: '100vh'
        }}>
            <Card
                className="redbook-card"
                title={
                    <div style={{
                        fontSize: '24px',
                        background: 'var(--background-gradient)',
                        WebkitBackgroundClip: 'text',
                        WebkitTextFillColor: 'transparent'
                    }}>编辑个人资料</div>
                }
                style={{ 
                    maxWidth: 800,
                    margin: '0 auto',
                    border: 'none'
                }}
            >
                <Form
                    layout="vertical"
                    initialValues={{
                        birthday: moment(data.Birthday, 'YYYY-MM-DD'),
                        nickname: data.Nickname,
                        aboutMe: data.AboutMe
                    }}
                    onFinish={onFinish}
                    onFinishFailed={onFinishFailed}
                >
                    <Form.Item
                        label="昵称"
                        name="nickname"
                    >
                        <Input className="redbook-input" />
                    </Form.Item>

                    <Form.Item
                        label="生日"
                        name="birthday"
                    >
                        <DatePicker 
                            className="redbook-input"
                            style={{ width: '100%' }}
                            format="YYYY-MM-DD"
                        />
                    </Form.Item>

                    <Form.Item
                        label="关于我"
                        name="aboutMe"
                    >
                        <TextArea 
                            style={{ 
                                borderRadius: '8px',
                                padding: '12px'
                            }}
                            rows={6}
                            placeholder="分享你的故事..."
                        />
                    </Form.Item>

                    <Form.Item>
                        <Button 
                            className="redbook-button"
                            type="primary" 
                            htmlType="submit"
                            block
                        >
                            保存修改
                        </Button>
                    </Form.Item>
                </Form>
            </Card>
        </div>
    );
}

export default EditForm;