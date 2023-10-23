<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;
use Illuminate\Support\Facades\Queue;

/**
 * @property int $id
 * @property array $assignees
 * @property TaskType $type
 */
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
        'completed',
        'type',
    ];

    protected $casts = [
        'dueDate' => 'datetime',
        'completedAt' => 'datetime',
    ];

    protected function getAssigneesAttribute(): array
    {
        return $this->taskAssignees->pluck('id')->toArray();
    }

    protected function setAssigneesAttribute(array $assignees): void
    {
        $this->taskAssignees()->sync($assignees);
    }

    public function taskAssignees(): BelongsToMany
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

    protected function getTypeAttribute(): TaskType
    {
        return $this->taskType;
    }

    protected function setTypeAttribute($value): void
    {
        $this->taskType()->associate($value);
    }

    public function taskType(): BelongsTo
    {
        return $this->belongsTo(TaskType::class, 'task_type_id');
    }
}
