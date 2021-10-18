import React from 'react';

interface Props {
  comments: string[]
}

const Textarea = ({ comments }: Props)  => {
  return (
    <>
      <textarea value={comments.join('\n')} rows={10} readOnly></textarea>
    </>
  );
};

export default Textarea;