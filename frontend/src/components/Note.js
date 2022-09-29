import React from 'react'
import { MdDeleteForever } from 'react-icons/md'

function Note({ id, text, date, time, handleDeleteNote}) {
  return (
    <div className='note'>
        <span>{text}</span>
        <div className='note-footer'>
            <small>{date} | {time}</small>
            <MdDeleteForever onClick={() => handleDeleteNote(id)}className='delete-icon' size='1.3em' />
        </div>
    </div>
  )
}

export default Note