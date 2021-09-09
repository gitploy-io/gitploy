import { Form, Input, Button, Space } from "antd"
import { Repo } from "../models"

export interface RepoSettingsFormProps {
    repo: Repo
    saving: boolean
    onClickSave(payload: {configPath: string}): void
    onClickLock(): void
    onClickUnlock(): void
    onClickDeactivate(): void
}

export default function RepoSettingForm(props: RepoSettingsFormProps): JSX.Element {
    const layout = {
      labelCol: { span: 5},
      wrapperCol: { span: 12 },
    };

    const submitLayout = {
      wrapperCol: { offset: 5, span: 12 },
    };

    const onFinish = (values: any) => {
        const payload = {
            configPath: values.config
        }
        props.onClickSave(payload)
    }

    const values = {
        "config": props.repo.configPath
    }

    return (
        <Form
            name="setting"
            initialValues={values}
            onFinish={onFinish}
        >
            <Form.Item
                label="Config"
                name="config"
                rules={[{required: true}]}
                {...layout}>
                <Input />
            </Form.Item>
            <Form.Item {...submitLayout}>
                <Form.Item noStyle>
                    <Button 
                        loading={props.saving}
                        type="primary" 
                        htmlType="submit">
                      Submit
                    </Button>
                </Form.Item>
            </Form.Item>
            <Form.Item {...submitLayout}>
                <Form.Item noStyle>
                    <Space>
                        {(props.repo.locked)? 
                            <Button
                                danger
                                type="primary"
                                onClick={props.onClickUnlock}
                            >
                                UNLOCK REPOSITORY
                            </Button>:
                            <Button
                                danger
                                onClick={props.onClickLock}
                            >
                                LOCK REPOSITORY
                            </Button>
                        }
                        <Button
                            danger
                            type="primary"
                            onClick={props.onClickDeactivate}
                        >
                            DEACTIVATE
                        </Button>
                    </Space>
                </Form.Item>
            </Form.Item>
        </Form>
    )
}