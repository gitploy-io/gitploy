import { Form, Input, Button, Space, Typography } from "antd"
import { Repo } from "../models"

export interface RepoSettingsFormProps {
    repo: Repo
    saving: boolean
    onClickSave(payload: {configPath: string}): void
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
                {...layout}
            >
                <Space>
                    <Form.Item
                        name="config"
                        rules={[{required: true}]}
                        noStyle
                    >
                        <Input />
                    </Form.Item>
                    <Typography.Link target="_blank" href={`/link/${props.repo.namespace}/${props.repo.name}/config`}>
                        Link
                    </Typography.Link>
                </Space>
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