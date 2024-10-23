// import data from "../../../../data.json";
// import React, {useState} from "react";
//
// export default function Durations({ item }) {
//     const [activeIndex, setActiveIndex] = useState(0);
//
//     const handleDurationClick = (index) => {
//         setActiveIndex(index === activeIndex ? null : index)
//     }
//
//     return (
//         {Array.from({length: Object.keys(data[item].costs).length}, (_, i) => (
//             <div className={`modal__duration flex bgc-1 b br10 ${activeIndex === i ? 'duration-active' : ''}`} key={i} onClick={() => handleDurationClick(i)}>
//                 <div className={`modal__duration-checkbox`}></div>
//                 <div className={'modal__duration-text flex'}>
//                     <p>{ data[item].costs[`${i + 1}`][1] }</p>
//                     <span>/</span>
//                     <p>{ data[item].costs[`${i + 1}`][0] }</p>
//                 </div>
//             </div>
//             )
//         )}
//     )
// }