import { Modal, Form, Select, Input, InputNumber, Checkbox } from 'antd';

import {
  Env,
  DynamicPayloadInput,
  DynamicPayloadInputTypeEnum,
} from '../../models';

export interface DynamicPayloadModalProps {
  visible: boolean;
  env: Env;
  onClickOk(values: any): void;
  onClickCancel(): void;
}

export default function DynamicPayloadModal(
  props: DynamicPayloadModalProps
): JSX.Element {
  const [form] = Form.useForm();

  const onClickOk = () => {
    form
      .validateFields()
      .then((values) => {
        props.onClickOk(values);
      })
      .catch((info) => {
        console.log(info);
      });
  };

  const onClickCancel = () => {
    props.onClickCancel();
  };

  // Build items dynamically
  const items = new Array<JSX.Element>();
  if (props.env.dynamicPayload) {
    // Object.entries(props.env.dynamicPayload.inputs).forEach()
    Object.entries(props.env.dynamicPayload.inputs).forEach((entry) => {
      const [name, input] = entry;
      items.push(<DynamicItem key={name} name={name} input={input} />);
    });
  }

  // Build the initialValues
  const initialValues: any = {};
  if (props.env.dynamicPayload) {
    Object.entries(props.env.dynamicPayload.inputs).forEach((entry) => {
      const [name, input] = entry;
      initialValues[name] = input.default;
    });
  }

  return (
    <Modal visible={props.visible} onOk={onClickOk} onCancel={onClickCancel}>
      <Form
        form={form}
        layout="vertical"
        name="dynamic_payload"
        initialValues={initialValues}
      >
        {items.map((item) => item)}
      </Form>
    </Modal>
  );
}

interface DynamicItemProps {
  name: string;
  input: DynamicPayloadInput;
}

function DynamicItem({ name, input }: DynamicItemProps): JSX.Element {
  // Capitalize the first character.
  const label = name.charAt(0).toUpperCase() + name.slice(1);
  const description = input.description;
  const rules = input.required ? [{ required: true }] : [];

  switch (input.type) {
    case DynamicPayloadInputTypeEnum.Select:
      return (
        <Form.Item
          label={label}
          name={name}
          tooltip={description}
          rules={rules}
        >
          <Select>
            {input.options?.map((option: any, idx: any) => (
              <Select.Option key={idx} value={option}>
                {option}
              </Select.Option>
            ))}
          </Select>
        </Form.Item>
      );
    case DynamicPayloadInputTypeEnum.String:
      return (
        <Form.Item
          label={label}
          name={name}
          tooltip={description}
          rules={rules}
        >
          <Input />
        </Form.Item>
      );
    case DynamicPayloadInputTypeEnum.Number:
      return (
        <Form.Item
          label={label}
          name={name}
          tooltip={description}
          rules={rules}
        >
          <InputNumber />
        </Form.Item>
      );
    case DynamicPayloadInputTypeEnum.Boolean:
      return (
        <Form.Item
          label={label}
          name={name}
          tooltip={description}
          rules={rules}
        >
          <Checkbox />
        </Form.Item>
      );
  }
}
