import React from 'react'

function Header({ handleToggleDarkMode }) {
  return (
    <div className='header'>
        <h1>Notes</h1>
        <button className='save' onClick={() => handleToggleDarkMode((previousDarkMode)=> !previousDarkMode)}>Toggle Mode</button>
    </div>
  )
}

export default Header