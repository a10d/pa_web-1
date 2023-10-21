<?php

namespace App\Http\Requests;

use Illuminate\Contracts\Validation\ValidationRule;
use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Support\Facades\Auth;

class UpdateTask extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return Auth::check()
            && Auth::user()->can('update', [$this->task()]);
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

            'assignees' => ['array', 'required'],
            'assignees.*' => ['integer', 'exists:users,id'],

            'description' => ['string', 'required'],
            'dueDate' => ['date', 'required', 'after:today'],

            'completed' => ['boolean', 'required'],
        ];
    }
}
