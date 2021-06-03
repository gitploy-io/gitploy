import { Button } from 'antd'

interface PaginationProps {
    page: number
    isLast: boolean
    onClickPrev(): void
    onClickNext(): void
}

export default function Pagination(props: PaginationProps) {
    const isFirst = props.page <= 1

    return (
        <div>
            <Button disabled={isFirst} onClick={props.onClickPrev} style={{borderTopRightRadius: 0, borderBottomRightRadius: 0}}>Prev</Button>
            <Button disabled={props.isLast} onClick={props.onClickNext} style={{borderTopLeftRadius: 0, borderBottomLeftRadius: 0}}>Next</Button>
        </div>
    )
}