import { useState } from 'react';
import { Button, Typography, List, Avatar } from 'antd';
import moment from 'moment';

import { Commit } from '../../models';

interface CommitChangesProps {
  changes: Commit[];
}

export default function CommitChanges(props: CommitChangesProps): JSX.Element {
  return (
    <List
      style={{
        maxHeight: 400,
        overflow: 'auto',
      }}
      bordered
      dataSource={props.changes}
      renderItem={(commit, idx) => {
        return <CommitChange key={idx} commit={commit} />;
      }}
    />
  );
}

function CommitChange(props: { commit: Commit }): JSX.Element {
  const [message, ...description] = props.commit.message.split(/(\r\n|\n|\r)/g);

  const [hide, setHide] = useState(true);

  const onClickHide = () => {
    setHide(!hide);
  };

  return (
    <List.Item>
      <List.Item.Meta
        title={
          <>
            <a href={props.commit.htmlUrl} target="_blank">
              {message}
            </a>
            {/* Display the description when the button is clicked. */}
            {description.length ? (
              <Button size="small" type="text" onClick={onClickHide}>
                <Typography.Text className="gitploy-code" code>
                  ...
                </Typography.Text>
              </Button>
            ) : (
              <></>
            )}
            {!hide ? (
              <Typography.Paragraph style={{ margin: 0 }}>
                <pre style={{ marginBottom: 0, fontSize: 12 }}>
                  {description.join('').trim()}
                </pre>
              </Typography.Paragraph>
            ) : (
              <></>
            )}
          </>
        }
        description={
          props.commit?.author ? (
            <>
              <Avatar size="small" src={props.commit.author.avatarUrl} />
              &nbsp;
              <Typography.Text strong>
                {props.commit.author.login}
              </Typography.Text>{' '}
              committed&nbsp;
              {moment(props.commit.author?.date).fromNow()}
            </>
          ) : (
            <></>
          )
        }
      />
      <div style={{ marginLeft: 30 }}>
        <Button href={props.commit.htmlUrl} target="_blank">
          {props.commit.sha.substring(0, 7)}
        </Button>
      </div>
    </List.Item>
  );
}
