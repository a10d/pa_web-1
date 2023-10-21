<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;

class Task extends Model
{
    use HasFactory;


    protected $fillable = [
        'type',
        'title',
        'description',
        'dueDate',
        'assignees',
        'completed',
    ];

    protected $appends = [
        'completed'
    ];

    public function type(): BelongsTo
    {
        return $this->belongsTo(TaskType::class, 'task_type_id');
    }

    public function assignees(): BelongsToMany
    {
        return $this->belongsToMany(User::class, 'task_assignees', 'task_id', 'user_id');
    }

    protected function getCompletedAttribute(): bool
    {
        return $this->completedAt !== null;
    }

    protected function setCompletedAttribute(bool $completed): void
    {
        $this->attributes['completedAt'] = $completed ? ($this->attributes['completedAt'] ?? now()) : null;
    }
}
