import React from "react";

export function calculateTimeLeft(endTime) {
    const currentTime = new Date() / 1000;
    if (endTime === 0) {
        return <span className={'br10 pd-8 forever'}>Навсегда</span>

    }
    if (currentTime > endTime) {
        return <span className={'br10 pd-8 elapsed'}>Срок истек</span>
    }
    const hours = Math.floor(((endTime - currentTime) / 60) / 60)
    const minutes = Math.floor(((endTime - currentTime) / 60) % 60)
    return `${hours} час(а/ов) ${minutes} минут(ы)`
}

export function calculateTimeDescription (startTime, endTime) {
    const start = new Date(startTime * 1000);
    const end = new Date(endTime * 1000);
    const bannedAt = `${start.toDateString()} ${start.getHours().toString().padStart(2, '0')}:${start.getMinutes().toString().padStart(2, '0')}`;

    let duration

    if (endTime === 0) {
        duration = 'Без срока окончания'
    } else {
        const durationMs = end - start;

        const durationDays = Math.floor(durationMs / (1000 * 60 * 60 * 24));
        const durationHours = Math.floor((durationMs % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
        const durationMinutes = Math.floor((durationMs % (1000 * 60 * 60)) / (1000 * 60));

        duration = `${durationDays} дн., ${durationHours} час., ${durationMinutes} мин.`;
    }

    return `• Время выдачи: ${bannedAt}\n• Выдан на: ${duration}`
}