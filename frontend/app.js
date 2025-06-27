function TokenViewer() {
  const [tokens, setTokens] = React.useState([]);
  const [selection, setSelection] = React.useState({start: null, end: null});
  const [isSelecting, setIsSelecting] = React.useState(false);

  React.useEffect(() => {
    fetch('/content')
      .then(res => res.json())
      .then(data => {
        if (data && Array.isArray(data.tokens)) {
          setTokens(data.tokens);
        }
      });
  }, []);

  const handleMouseDown = index => {
    setIsSelecting(true);
    setSelection({start: index, end: index});
  };

  const handleMouseEnter = index => {
    if (isSelecting) {
      setSelection(sel => ({...sel, end: index}));
    }
  };

  const handleMouseUp = () => {
    setIsSelecting(false);
  };

  const isHighlighted = index => {
    const {start, end} = selection;
    if (start === null || end === null) return false;
    const [s, e] = start < end ? [start, end] : [end, start];
    return index >= s && index <= e;
  };

  return (
    <div className="text" onMouseUp={handleMouseUp}>
      {tokens.map((tok, idx) => (
        <span
          key={idx}
          className={'token' + (isHighlighted(idx) ? ' highlight' : '')}
          onMouseDown={() => handleMouseDown(idx)}
          onMouseEnter={() => handleMouseEnter(idx)}
        >
          {tok}
        </span>
      ))}
    </div>
  );
}

ReactDOM.render(<TokenViewer />, document.getElementById('root'));

