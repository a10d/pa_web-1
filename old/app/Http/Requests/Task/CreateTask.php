<?php

namespace App\Http\Requests\Task;

use App\Models\Task;
use Illuminate\Contracts\Validation\ValidationRule;
use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Support\Facades\Auth;

class CreateTask extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return Auth::check()
            && Auth::user()->can('create', [Task::class]);
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'title' => ['string', 'required'],

            'type' => ['integer', 'required', 'exists:task_types,id'],

            'assignees' => ['array', 'nullable'],
            'assignees.*' => ['integer', 'exists:users,id'],

            'description' => ['string', 'required'],
            'dueDate' => ['date', 'required', 'after:today'],
        ];
    }
}
