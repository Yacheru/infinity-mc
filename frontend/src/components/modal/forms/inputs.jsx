export default function Inputs({ item, handleChange, colorValue, badgeValue }) {
    switch (item) {
        case 'nickname':
            return (
                <input className={'modal__input b'} type="text" id={'color'} value={colorValue} placeholder={'Укажите желаемый цвет в HEX формате'} onChange={handleChange}/>
            );
        case 'badge':
            return (
                <div>
                    <input className={'modal__input b'} type="text" id={'color'} value={colorValue} placeholder={'Укажите желаемый цвет значка в HEX формате'} onChange={handleChange}/>
                    <input className={'modal__input b'} type="text" id={'badge'} value={badgeValue} placeholder={'Укажите желаемый значок из UTF-8'} onChange={handleChange}/>
                </div>
            );
    }
}