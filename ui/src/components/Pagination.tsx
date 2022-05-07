import { Button } from 'antd';

export interface PaginationProps {
  disabledPrev: boolean;
  disabledNext: boolean;
  onClickPrev(): void;
  onClickNext(): void;
}

export default function Pagination(props: PaginationProps): JSX.Element {
  return (
    <div>
      <Button
        disabled={props.disabledPrev}
        onClick={props.onClickPrev}
        style={{ borderTopRightRadius: 0, borderBottomRightRadius: 0 }}
      >
        Prev
      </Button>
      <Button
        disabled={props.disabledNext}
        onClick={props.onClickNext}
        style={{ borderTopLeftRadius: 0, borderBottomLeftRadius: 0 }}
      >
        Next
      </Button>
    </div>
  );
}
