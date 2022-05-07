import { Form, Input, Button, Space, Typography } from 'antd';
import { useState } from 'react';

export interface SettingFormProps {
  configLink: string;
  initialValues?: SettingFormValues;
  onClickFinish(values: SettingFormValues): void;
  onClickDeactivate(): void;
}

export interface SettingFormValues {
  name: string;
  config_path: string;
}

export default function SettingForm({
  configLink,
  initialValues,
  onClickFinish,
  onClickDeactivate,
}: SettingFormProps): JSX.Element {
  const [saving, setSaving] = useState(false);

  const layout = {
    labelCol: { span: 5 },
    wrapperCol: { span: 12 },
  };

  const submitLayout = {
    wrapperCol: { offset: 5, span: 12 },
  };

  const onFinish = (values: any) => {
    setSaving(true);
    onClickFinish(values);
    setSaving(false);
  };

  return (
    <Form name="setting" initialValues={initialValues} onFinish={onFinish}>
      <Form.Item
        label="Name"
        name="name"
        {...layout}
        rules={[{ required: true }]}
      >
        <Input />
      </Form.Item>
      <Form.Item label="Config" {...layout}>
        <Space>
          <Form.Item name="config_path" rules={[{ required: true }]} noStyle>
            <Input />
          </Form.Item>
          <Typography.Link target="_blank" href={configLink}>
            Link
          </Typography.Link>
        </Space>
      </Form.Item>
      <Form.Item {...submitLayout}>
        <Form.Item noStyle>
          <Button loading={saving} type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form.Item>
      <Form.Item {...submitLayout}>
        <Form.Item noStyle>
          <Space>
            <Button danger type="primary" onClick={onClickDeactivate}>
              DEACTIVATE
            </Button>
          </Space>
        </Form.Item>
      </Form.Item>
    </Form>
  );
}
