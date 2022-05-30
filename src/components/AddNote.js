 import { useState } from 'react';
 
 function AddNote({ handleAddNote }) {
    const [noteText, setNoteText] = useState('');
    const characterLimit = 200;

    const handleChange = (event) => {
        if(noteText.length <= characterLimit){
            setNoteText(event.target.value);
        }
    }

    const handleSaveClick = () => {
        if(noteText.trim().length > 0){
            handleAddNote(noteText)
            setNoteText('')
        }
    }

    return (
        <>
            <div className='note new'>
                <textarea
                    rows='8'
                    cols='10'
                    placeholder='I am grateful for.. &#10; '
                    value={noteText}
                    onChange={handleChange}
                ></textarea>
                    <div className='note-footer'>
                        <small> {characterLimit - noteText.length} characters remaining</small>
                        <button className='save' onClick={handleSaveClick}>Save</button>
                    </div>
            </div>
        </>
    )
 }
 
 export default AddNote