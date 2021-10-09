import React from 'react';

interface Props {
  value: string[]
}

const Textarea = ({ value }: Props)  => {
  return (
    <>
      <textarea value={value.join('\n')} rows={10} readOnly></textarea>
    </>
  );
};

export default Textarea;