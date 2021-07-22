import { Spin as S} from 'antd';
import { LoadingOutlined } from '@ant-design/icons';

const antIcon = <LoadingOutlined style={{ fontSize: 24 }} spin />;

export default function Spin(): JSX.Element {
    return (
        <S indicator={antIcon} />
    )
}