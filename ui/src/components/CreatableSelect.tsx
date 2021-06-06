import { useState } from 'react';
import { Select, Divider, Input } from 'antd'
import { PlusOutlined } from '@ant-design/icons';

interface CreatableSelectProps {
    options: Option[]
    onSelectOption(option: Option): void
    onClickAddItem(option: Option): void
    style?: React.CSSProperties
    placeholder?: string
}

export interface Option {
    label: string
    value: string
}

export default function CreatableSelect(props: CreatableSelectProps) {
    const initOption = {label: "", value: ""}
    const [item, setItem] = useState<Option>(initOption)

    const onChangeItem = (e: any) => {
        const value = e.target.value
        setItem({
            label: value,
            value: value
        })
    }

    const onClickAddItem = () => {
        props.onClickAddItem(item)
        setItem(initOption)
    }

    const onSelectOption = (value: string) => {
        const option = props.options.find(o => o.value === value)

        if (option === undefined) throw new Error("The option doesn't exist.")

        props.onSelectOption(option)
    }

    return (
        <Select
            onSelect={onSelectOption}
            style={{...props.style}}
            placeholder={props.placeholder}
            dropdownRender={menu => (
            <div>
                {menu}
                <Divider style={{ margin: '4px 0' }} />
                <div style={{ display: 'flex', flexWrap: 'nowrap' }}>
                    <Input 
                        value={item.value}
                        onChange={onChangeItem}
                        placeholder="Add item manually"
                        bordered={false} />
                    {/* eslint-disable-next-line jsx-a11y/anchor-is-valid */}
                    <a
                        style={{ flex: 'none', padding: '4px', display: 'block', cursor: 'pointer' }} 
                        onClick={onClickAddItem}>
                      <PlusOutlined /> Add item
                    </a>
                </div>
            </div>
        )}>
            {props.options.map((option, idx) => {
                return <Select.Option key={idx} value={option.value}>{option.label}</Select.Option>
            })}
        </Select>
    )
}