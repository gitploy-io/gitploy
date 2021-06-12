import { Form, Input, Button, Space } from "antd"
import { Repo, RepoPayload } from "../models"

export interface SettingsFormProps {
    repo: Repo
    saving: boolean
    onClickSave(payload:RepoPayload): void
    onClickDeactivate(): void
}

export default function SettingForm(props: SettingsFormProps) {
    const layout = {
      labelCol: { span: 5},
      wrapperCol: { span: 12 },
    };

    const submitLayout = {
      wrapperCol: { offset: 5, span: 12 },
    };

    const onFinish = (values: any) => {
        const payload: RepoPayload = {
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
            onFinish={onFinish}>
            <Form.Item
                label="Config"
                name="config"
                rules={[{required: true}]}
                {...layout}>
                <Input />
            </Form.Item>
            <Form.Item {...submitLayout}>
                <Space>
                    <Form.Item noStyle>
                        <Button 
                            loading={props.saving}
                            type="primary" 
                            htmlType="submit">
                          Submit
                        </Button>
                    </Form.Item>
                    <Button
                        danger
                        type="primary"
                        onClick={props.onClickDeactivate}>
                        DEACTIVATE
                    </Button>
                </Space>
            </Form.Item>
        </Form>
    )
}