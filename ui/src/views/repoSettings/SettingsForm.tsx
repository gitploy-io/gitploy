import { shallowEqual } from "react-redux";
import { Form, Input, Button, Space, Typography } from "antd"

import { useAppSelector, useAppDispatch } from "../../redux/hooks"
import { save, deactivate, repoSettingsSlice as slice } from "../../redux/repoSettings"
import { RequestStatus } from "../../models"

export default function RepoSettingForm(): JSX.Element {
    const { repo, saving } = useAppSelector(state => state.repoSettings, shallowEqual)
    const dispatch = useAppDispatch()

    const layout = {
      labelCol: { span: 5},
      wrapperCol: { span: 12 },
    };

    const submitLayout = {
      wrapperCol: { offset: 5, span: 12 },
    };

    const onClickFinish = (values: any) => {
        dispatch(slice.actions.setConfigPath(values.config))
        dispatch(save())
    }

    const onClickDeactivate = () => {
        dispatch(deactivate())
    }

    const initialValues = {
        "config": repo?.configPath
    }

    return (
        <Form
            name="setting"
            initialValues={initialValues}
            onFinish={onClickFinish}
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
                    <Typography.Link target="_blank" href={`/link/${repo?.namespace}/${repo?.name}/config`}>
                        Link
                    </Typography.Link>
                </Space>
            </Form.Item>
            <Form.Item {...submitLayout}>
                <Form.Item noStyle>
                    <Button 
                        loading={saving === RequestStatus.Pending}
                        type="primary" 
                        htmlType="submit"
                    >
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
                            onClick={onClickDeactivate}
                        >
                            DEACTIVATE
                        </Button>
                    </Space>
                </Form.Item>
            </Form.Item>
        </Form>
    )
}