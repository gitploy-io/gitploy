import { useState } from 'react';
import { Select, Divider, Input, SelectProps } from 'antd'
import { PlusOutlined } from '@ant-design/icons';

interface CreatableSelectProps extends SelectProps<string>{
    options: Option[]
    onSelectOption(option: Option): void
    onClickAddItem(option: Option): void
}

export interface Option {
    label: string
    value: string
}

export default function CreatableSelect(props: CreatableSelectProps) {
    const initOption = {label: "", value: ""}
    const [item, setItem] = useState<Option>(initOption)

    // Clone Select props only
    // https://stackoverflow.com/questions/34698905/how-can-i-clone-a-javascript-object-except-for-one-key
    const {options, onSelectOption, onClickAddItem, ...selectProps} = props

    const _onChangeItem = (e: any) => {
        const value = e.target.value
        setItem({
            label: value,
            value: value
        })
    }

    const _onClickAddItem = () => {
        onClickAddItem(item)
        setItem(initOption)
    }

    const _onSelectOption = (value: string) => {
        const option = props.options.find(o => o.value === value)

        if (option === undefined) throw new Error("The option doesn't exist.")

        onSelectOption(option)
    }

    return (
        <Select
            {...selectProps}
            onSelect={_onSelectOption}
            dropdownRender={menu => (
            <div>
                {menu}
                <Divider style={{ margin: '4px 0' }} />
                <div style={{ display: 'flex', flexWrap: 'nowrap' }}>
                    <Input 
                        value={item.value}
                        onChange={_onChangeItem}
                        placeholder="Add item manually"
                        bordered={false} />
                    {/* eslint-disable-next-line jsx-a11y/anchor-is-valid */}
                    <a
                        style={{ flex: 'none', padding: '4px', display: 'block', cursor: 'pointer' }} 
                        onClick={_onClickAddItem}>
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